package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
)

var fileSrc = `package main
import "fmt"

func main() {
	fmt.Println("Hello")
}
`

type astVisitor func(ast.Node) bool

func (v astVisitor) Visit(node ast.Node) ast.Visitor {
	if v(node) {
		return v
	}
	return nil
}

func main() {
	fmt.Println("Analyzer...")

	fset := token.NewFileSet()

	f, err := parser.ParseFile(fset, "main.go", fileSrc, parser.ParseComments)
	if err != nil {
		fmt.Println(err)
		return
	}

	var bf bytes.Buffer
	ast.Fprint(&bf, fset, f, func(string, reflect.Value) bool {
		return true
	})

	fmt.Println(bf.String())

	/*
		fn := func(node ast.Node) bool {
			fmt.Printf("%T\n", node)
			return true
		}
		ast.Walk(astVisitor(fn), f)
	*/
}
