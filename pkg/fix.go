/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package pkg

import (
	"fmt"
	"go/ast"
	"io"
	"os"
	"strings"
	"time"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"golang.org/x/tools/go/ast/astutil"
	"golang.org/x/tools/go/packages"
)

// FixOptions controls the behavior of a fix operation
type FixOptions struct {
	WriteOnError bool
	Overwrite    bool
	Verbose      bool
	Packages     []string
	Dir          string
	Out          io.Writer
}

// DefaultFixOptions returns a default set of options
func DefaultFixOptions() *FixOptions {
	return &FixOptions{
		WriteOnError: false,
		Overwrite:    true,
		Verbose:      true,
		Out:          os.Stdout,
	}
}

// Complete populates any derived options
func (o *FixOptions) Complete() error {
	if o.Out == nil {
		o.Out = os.Stdout
	}
	if len(o.Packages) == 0 {
		o.Packages = []string{"."}
	}
	return nil
}

// Validate ensures the specified options are valid
func (o *FixOptions) Validate() error {
	if o.Out == nil {
		return fmt.Errorf("no output specified")
	}
	if len(o.Packages) == 0 {
		return fmt.Errorf("no packages specified")
	}
	return nil
}

// Run executes the fix operation
func (o *FixOptions) Run() error {
	start := time.Now()
	pkgs, err := packages.Load(&packages.Config{
		Mode:  packages.NeedName | packages.NeedFiles | packages.NeedCompiledGoFiles | packages.NeedTypes | packages.NeedTypesInfo | packages.NeedSyntax | packages.NeedDeps,
		Tests: true,
		Dir:   o.Dir,
	}, o.Packages...)
	if err != nil {
		return fmt.Errorf("error loading packages: %w", err)
	}

	if o.Verbose {
		if len(pkgs) == 1 {
			fmt.Fprintf(o.Out, "loaded %d package in %v\n", len(pkgs), time.Now().Sub(start))
		} else {
			fmt.Fprintf(o.Out, "loaded %d packages in %v\n", len(pkgs), time.Now().Sub(start))
		}
	}

	errorsEncountered := 0
	processed := map[string]bool{}
	for _, pkg := range pkgs {
		dec := decorator.NewDecorator(pkg.Fset)
		for i, filename := range pkg.CompiledGoFiles {
			if !strings.HasSuffix(filename, ".go") {
				continue
			}
			if processed[filename] {
				// the same file can appear in test and non-test packages, only process once
				continue
			}
			processed[filename] = true

			// Get the decorated file to preserve comment association with nodes before we start moving them around
			decFile, err := dec.DecorateFile(pkg.Syntax[i])
			if err != nil {
				return fmt.Errorf("error processing %s: %w", filename, err)
			}

			transformContext := &transformFileContext{
				dec:      dec,
				decFile:  decFile,
				pkg:      pkg,
				file:     pkg.Syntax[i],
				modified: false,
				verbose:  o.Verbose,
				out:      o.Out,
			}

			// Walk the AST, applying transforms to function calls
			astutil.Apply(pkg.Syntax[i],
				nil,
				func(c *astutil.Cursor) bool {
					callExpr, isCall := c.Node().(*ast.CallExpr)
					if !isCall {
						return true
					}
					funSelector, isSelector := callExpr.Fun.(*ast.SelectorExpr)
					if !isSelector {
						return true
					}
					selectorType := pkg.TypesInfo.TypeOf(funSelector.X)
					if selectorType == nil {
						return true
					}
					selectorTypeName := strings.TrimPrefix(selectorType.String(), "*")

					// update context for this cursor and call
					transformContext.cursor = c
					transformContext.call = callExpr
					transformContext.decCall = dec.Dst.Nodes[callExpr].(*dst.CallExpr)
					transformContext.function = funSelector.Sel.Name
					transformContext.pos = pkg.Fset.Position(callExpr.Pos())
					// iterate over transforms
					for _, transform := range transforms {
						if transform.matcher(selectorTypeName, funSelector.Sel.Name) {
							transform.transform.Transform(transformContext)
						}
					}
					return true
				},
			)

			// If a transform modified the AST, write the results
			if transformContext.modified {
				func() {
					// Only write with errors if they opted in
					if transformContext.err != nil && !o.WriteOnError {
						fmt.Fprintf(o.Out, "%serrors encountered, skipping write (run with --force to write anyway)\n", fileIndent)
						return
					}

					// Create a temp file
					tmpFilename := pkg.CompiledGoFiles[i] + ".tmp"
					f, err := os.Create(tmpFilename)
					defer f.Close()
					if err != nil {
						fmt.Fprintf(o.Out, "%s%v\n", fileIndent, err)
						return
					}

					// Write the AST
					err = decorator.Fprint(f, decFile)
					if err != nil {
						fmt.Fprintf(o.Out, "%s%v\n", fileIndent, err)
						return
					}

					// Move the file into place
					if o.Overwrite {
						err = os.Rename(tmpFilename, pkg.CompiledGoFiles[i])
						if err != nil {
							fmt.Fprintf(o.Out, "%s%v\n", fileIndent, err)
						}
					}
				}()
			}
			if transformContext.err != nil {
				errorsEncountered++
			}
		}
	}

	if errorsEncountered == 1 {
		return fmt.Errorf("errors encountered in %d file", errorsEncountered)
	} else if errorsEncountered > 1 {
		return fmt.Errorf("errors encountered in %d files", errorsEncountered)
	}

	return nil
}
