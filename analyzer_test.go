package main_test

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
	"testing"

	"github.com/a6cexz/goanalyzer/diag/syntax/syntaxkind"
	"github.com/stretchr/testify/assert"
)

func TestParseFileAndPrintAst(t *testing.T) {
	var fileSrc = `package main
import "fmt"

func main() {
	fmt.Println("Hello")
}
`

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

	assert.True(t, true)
}

func TestCollectAllNodes(t *testing.T) {
	src := `package main
var a = 10
var b = 20
`

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "main.go", src, parser.ParseComments)
	if err != nil {
		t.Error(err)
		return
	}

	nodes := syntaxkind.CollectAllNodes(f)
	fmt.Println(len(nodes))
}
