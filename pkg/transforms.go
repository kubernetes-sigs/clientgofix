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
	"go/parser"
	"go/token"
	"path"
	"strings"

	"github.com/dave/dst"
)

type Transformer interface {
	Transform(*TransformFileContext)
}
type TransformFunc func(*TransformFileContext)

func (f TransformFunc) Transform(c *TransformFileContext) {
	f(c)
}

type Transforms []Transformer

func (t Transforms) Transform(c *TransformFileContext) {
	for _, transformer := range t {
		transformer.Transform(c)
	}
}

type makeArgExprFunc func(*TransformFileContext) (ast.Expr, error)

func getOrCreateImport(c *TransformFileContext, importPackage, preferredAlias string) (alias string, err error) {
	// the import package will be quoted
	expectedLiteral := fmt.Sprintf("%q", importPackage)

	// find an existing import for this path
	for _, i := range c.file.Imports {
		if i.Path.Value == expectedLiteral {
			if i.Name != nil {
				// return import alias
				return i.Name.String(), nil
			}
			// no import alias, return the last path segment
			return path.Base(importPackage), nil
		}
	}

	localName := preferredAlias
	if len(localName) == 0 {
		localName = path.Base(importPackage)
	}
	// Scan existing import aliases/variable/func/type/const names in the file to avoid conflicts
	used := map[string]bool{}
	for _, i := range c.file.Imports {
		if i.Name != nil {
			// fmt.Println("import", i.Name.String())
			used[i.Name.String()] = true
		} else {
			// fmt.Println("import", path.Base(strings.Trim(i.Path.Value, `"`)))
			used[path.Base(strings.Trim(i.Path.Value, `"`))] = true
		}
	}
	ast.Inspect(c.file, func(n ast.Node) bool {
		if n == nil {
			return true
		}
		switch n := n.(type) {
		case *ast.FuncDecl:
			// fmt.Println("func", n.Name.Name)
			if n.Name != nil {
				used[n.Name.Name] = true
			}
		case *ast.ValueSpec:
			for _, name := range n.Names {
				// fmt.Println("value", name.Name)
				used[name.Name] = true
			}
		case *ast.TypeSpec:
			if n.Name != nil {
				// fmt.Println("type", n.Name.Name)
				used[n.Name.Name] = true
			}
		case *ast.AssignStmt:
			if n.Tok == token.DEFINE {
				for _, lhs := range n.Lhs {
					if ident, ok := lhs.(*ast.Ident); ok {
						// fmt.Printf("assign %v\n", ident.Name)
						used[ident.Name] = true
					}
				}
			}
		case *ast.Field:
			for _, name := range n.Names {
				// fmt.Println("field name", name.Name)
				used[name.Name] = true
			}
		default:
			// fmt.Printf("%T: %v\n", n, n)
		}
		return true
	})

	deconflicted := false
	if used[localName] {
		for i := 2; ; i++ {
			potentialName := fmt.Sprintf("%s%d", localName, i)
			if !used[potentialName] {
				localName = potentialName
				deconflicted = true
				break
			}
		}
	}

	// Create a new import spec
	importSpec := &ast.ImportSpec{Path: &ast.BasicLit{Kind: token.STRING, Value: expectedLiteral}}
	// Add an alias if specified
	if len(preferredAlias) > 0 || deconflicted {
		importSpec.Name = &ast.Ident{Name: localName}
	}

	// Create a decorated node
	decNode, err := c.dec.DecorateNode(importSpec)
	if err != nil {
		return "", err
	}
	decImportSpec := decNode.(*dst.ImportSpec)

	// Insert into first import declaration
	for _, decl := range c.file.Decls {
		astGenDecl, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}
		if astGenDecl.Tok != token.IMPORT {
			continue
		}
		dstGenDecl := c.dec.Dst.Nodes[astGenDecl].(*dst.GenDecl)

		// Append to imports
		c.file.Imports = append(c.file.Imports, importSpec)
		c.decFile.Imports = append(c.decFile.Imports, decImportSpec)

		// Insert into import declaration
		insert := 0
		for insert = 0; insert < len(astGenDecl.Specs); insert++ {
			neighborImportSpec, ok := astGenDecl.Specs[insert].(*ast.ImportSpec)
			if !ok {
				continue
			}
			if neighborImportSpec.Path.Value > expectedLiteral {
				break
			}
		}
		astGenDecl.Specs = append(
			astGenDecl.Specs[:insert],
			append([]ast.Spec{importSpec}, astGenDecl.Specs[insert:]...)...,
		)
		dstGenDecl.Specs = append(
			dstGenDecl.Specs[:insert],
			append([]dst.Spec{decImportSpec}, dstGenDecl.Specs[insert:]...)...,
		)

		c.Modified(fmt.Sprintf("added %s import", importPackage))
		return localName, nil
	}
	return "", fmt.Errorf("no import declaration block found")
}

func argToDecArg(c *TransformFileContext, arg ast.Expr, matchIndex int) (dst.Expr, error) {
	// Convert to decorated
	decoratedNode, err := c.dec.DecorateNode(arg)
	if err != nil {
		return nil, err
	}
	decArg, ok := decoratedNode.(dst.Expr)
	if !ok {
		return nil, fmt.Errorf("expected dst.Expr, got %T", decoratedNode)
	}

	// Match leading/trailing newlines of existing arg
	if matchIndex >= 0 && len(c.decCall.Args) > matchIndex {
		decArg.Decorations().Before = c.decCall.Args[matchIndex].Decorations().Before
		decArg.Decorations().After = c.decCall.Args[matchIndex].Decorations().After
	}

	return decArg, nil
}

// ensureArgAtIndex checks if the arg at the given position is of the specified type,
// and inserts a new arg made by f() at that position if it is not of the specified type.
func ensureArgAtIndex(pos int, argType string, f makeArgExprFunc) Transformer {
	return TransformFunc(func(c *TransformFileContext) {
		if len(c.call.Args) > pos {
			typeOf := c.pkg.TypesInfo.TypeOf(c.call.Args[pos])
			if typeOf == nil {
				// can happen on a rerun if there is a local variable masking the inserted package alias
				c.Warning(fmt.Sprintf("cannot determine type of arg %d, skipping inserting %s (check for variables shadowing import)", pos, argType))
				return
			}
			if typeOf.String() == argType {
				return
			}
		}
		// pos cannot be greater than the current length
		if pos > len(c.call.Args) {
			c.Error(fmt.Errorf("cannot set arg %d; current arg length is %d", pos, len(c.call.Args)))
			return
		}

		// Modify AST
		expr, err := f(c)
		if err != nil {
			c.Error(err)
			return
		}

		// Modify DST
		decArg, err := argToDecArg(c, expr, 0)
		if err != nil {
			c.Error(err)
			return
		}

		if pos == 0 {
			c.call.Args = append([]ast.Expr{expr}, c.call.Args...)
			c.decCall.Args = append([]dst.Expr{decArg}, c.decCall.Args...)
		} else if pos == len(c.call.Args) {
			c.call.Args = append(c.call.Args, expr)
			c.decCall.Args = append(c.decCall.Args, decArg)
		} else {
			c.call.Args = append(c.call.Args[:pos], append([]ast.Expr{expr}, c.call.Args[pos:]...)...)
			c.decCall.Args = append(c.decCall.Args[:pos], append([]dst.Expr{decArg}, c.decCall.Args[pos:]...)...)
		}

		c.Modified(fmt.Sprintf("inserted %s as arg %d", argType, pos))
	})
}
func ensureLastArg(argType string, f makeArgExprFunc) Transformer {
	return TransformFunc(func(c *TransformFileContext) {
		var msg string
		if len(c.call.Args) > 0 {
			typeOf := c.pkg.TypesInfo.TypeOf(c.call.Args[len(c.call.Args)-1])
			if typeOf == nil {
				msg = fmt.Sprintf("appended %s arg", argType)
			} else if typeOf.String() == argType {
				return
			} else {
				msg = fmt.Sprintf("last arg was %s, appended %s", typeOf.String(), argType)
			}
		} else {
			msg = fmt.Sprintf("appended %s arg", argType)
		}

		lastIndex := len(c.call.Args) - 1

		// Modify AST
		expr, err := f(c)
		if err != nil {
			c.Error(err)
			return
		}
		c.call.Args = append(c.call.Args, expr)

		// Modify DST
		decArg, err := argToDecArg(c, expr, lastIndex)
		if err != nil {
			c.Error(err)
			return
		}
		c.decCall.Args = append(c.decCall.Args, decArg)

		c.Modified(msg)
	})
}
func replaceArgAtIndexIfNil(argIndex int, argType string, f makeArgExprFunc) Transformer {
	return TransformFunc(func(c *TransformFileContext) {
		// argIndex cannot be greater than the current length
		if argIndex > len(c.call.Args) {
			// c.Error(fmt.Errorf("cannot set arg %d; current arg length is %d", argIndex, len(c.call.Args)))
			return
		}

		ident, isIdent := c.call.Args[argIndex].(*ast.Ident)
		if !isIdent {
			return
		}
		if ident.Name != "nil" {
			return
		}

		// Modify AST
		expr, err := f(c)
		if err != nil {
			c.Error(err)
			return
		}
		c.call.Args[argIndex] = expr

		// Modify DST
		decArg, err := argToDecArg(c, expr, argIndex)
		if err != nil {
			c.Error(err)
			return
		}
		c.decCall.Args[argIndex] = decArg

		c.Modified(fmt.Sprintf("replaced nil arg at %d with %s", argIndex, argType))
	})
}
func dereferenceArgAtIndexIfPointer(argIndex int, argType string) Transformer {
	return TransformFunc(func(c *TransformFileContext) {
		// argIndex cannot be greater than the current length
		if argIndex > len(c.call.Args) {
			// c.Error(fmt.Errorf("cannot set arg %d; current arg length is %d", argIndex, len(c.call.Args)))
			return
		}

		lastArg := c.call.Args[argIndex]
		typeOf := c.pkg.TypesInfo.TypeOf(lastArg)
		if typeOf == nil {
			return
		}
		if typeOf.String() != argType {
			return
		}

		if unary, isUnary := lastArg.(*ast.UnaryExpr); isUnary {
			if unary.Op == token.AND {
				if _, isComposite := unary.X.(*ast.CompositeLit); isComposite {
					// Modify AST
					c.call.Args[argIndex] = unary.X
					// Modify DST
					c.decCall.Args[argIndex] = c.decCall.Args[argIndex].(*dst.UnaryExpr).X
					c.Modified(fmt.Sprintf("replaced literal declaration of %s", argType))
					return
				}
			}
		}

		// Modify AST
		arg := &ast.UnaryExpr{Op: token.MUL, X: lastArg}
		c.call.Args[argIndex] = arg

		// Modify DST
		decArg, err := c.dec.DecorateNode(arg)
		if err != nil {
			c.Error(err)
			return
		}
		c.decCall.Args[argIndex] = decArg.(dst.Expr)

		c.Modified(fmt.Sprintf("dereferenced pointer arg at %d to %s", argIndex, argType))
	})
}

func makeContextArg(c *TransformFileContext) (ast.Expr, error) {
	contextAlias, err := getOrCreateImport(c, "context", "")
	if err != nil {
		return nil, err
	}
	return parser.ParseExpr(fmt.Sprintf("%s.TODO()", contextAlias))
}
func makeMetav1OptionsArg(optionsType string) makeArgExprFunc {
	return func(c *TransformFileContext) (ast.Expr, error) {
		metav1Alias, err := getOrCreateImport(c, "k8s.io/apimachinery/pkg/apis/meta/v1", "metav1")
		if err != nil {
			return nil, err
		}
		return parser.ParseExpr(fmt.Sprintf("%s.%s{}", metav1Alias, optionsType))
	}
}
