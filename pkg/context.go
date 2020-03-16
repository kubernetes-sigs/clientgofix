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
	"go/token"
	"io"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"golang.org/x/tools/go/ast/astutil"
	"golang.org/x/tools/go/packages"
)

var fileIndent = "  "

type TransformFileContext struct {
	verbose bool
	out     io.Writer

	dec     *decorator.Decorator
	decFile *dst.File
	decCall *dst.CallExpr

	cursor   *astutil.Cursor
	pkg      *packages.Package
	file     *ast.File
	call     *ast.CallExpr
	pos      token.Position
	function string

	err      error
	modified bool
	warned   bool
}

func (c *TransformFileContext) Warning(msg string) {
	if !c.modified && c.err == nil && !c.warned {
		fmt.Fprintf(c.out, "%v\n", c.pos.Filename)
	}
	fmt.Fprintf(c.out, "%s%v: %s: WARNING: %s\n", fileIndent, c.pos.Line, c.function, msg)
	c.warned = true
}
func (c *TransformFileContext) Error(err error) {
	if !c.modified && c.err == nil && !c.warned {
		fmt.Fprintf(c.out, "%v\n", c.pos.Filename)
	}
	fmt.Fprintf(c.out, "%s%v: %s: ERROR: %v\n", fileIndent, c.pos.Line, c.function, err)
	c.err = err
}
func (c *TransformFileContext) Modified(msg string) {
	if !c.modified && c.err == nil && !c.warned {
		fmt.Fprintf(c.out, "%v\n", c.pos.Filename)
	}
	if c.verbose {
		fmt.Fprintf(c.out, "%s%v: %s: %v\n", fileIndent, c.pos.Line, c.function, msg)
	}
	c.modified = true
}
