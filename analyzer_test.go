package main_test

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
	"testing"

	"golang.org/x/tools/go/ast/astutil"

	"github.com/a6cexz/goanalyzer/diag/text"
	"github.com/a6cexz/goanalyzer/diag/text/helpers"
	"github.com/stretchr/testify/assert"
)

type astVisitor func(ast.Node) bool

func (v astVisitor) Visit(node ast.Node) ast.Visitor {
	if v(node) {
		return v
	}
	return nil
}

func isExpr(node ast.Node) bool {
	expr := asExpr(node)
	return expr != nil
}
func asExpr(node ast.Node) ast.Expr {
	if node == nil {
		return nil
	}

	r, ok := node.(ast.Expr)
	if !ok {
		return nil
	}

	return r
}

func collectAllNodes(root ast.Node) []ast.Node {
	nodes := []ast.Node{}
	v := func(node ast.Node) bool {
		if node != nil {
			nodes = append(nodes, node)
		}
		return true
	}
	ast.Walk(astVisitor(v), root)
	return nodes
}

func getNodeAt(pos int) ast.Node {
	return nil
}

func parseTestFile(src string) (*token.FileSet, *ast.File, text.TextSpan, error) {
	src, m := helpers.RemoveTextMarkers(src)
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "source.go", src, parser.ParseComments)
	if err != nil {
		return fset, file, text.NewTextSpan(0, 0), err
	}

	start := 0
	if pos, ok := m["#start#"]; ok {
		start = pos
	}

	end := start
	if pos, ok := m["#end#"]; ok {
		end = pos
	}

	s := text.NewTextSpanFromBounds(start, end)
	return fset, file, s, nil
}

func getTestAstNode(src string) (ast.Node, error) {
	_, file, span, err := parseTestFile(src)
	if err != nil {
		return nil, err
	}

	start := token.Pos(span.Start())
	end := token.Pos(span.End())
	path, _ := astutil.PathEnclosingInterval(file, start, end)
	node := path[0]

	return node, nil
}

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

	nodes := collectAllNodes(f)
	fmt.Println(len(nodes))
}

func TestParseTestFile(t *testing.T) {
	src := `package main
var a = 1#start#0
`
	node, err := getTestAstNode(src)
	assert.NoError(t, err)
	assert.True(t, isExpr(node))
	assert.True(t, asExpr(node) != nil)
}
