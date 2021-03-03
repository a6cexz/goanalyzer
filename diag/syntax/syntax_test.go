package syntax_test

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"maxmapps.com/goanalyzer/diag/syntax"
)

func TestCommentGroupNode(t *testing.T) {
	c1 := &ast.Comment{
		Slash: token.Pos(1),
		Text:  "Test",
	}

	cg := &ast.CommentGroup{
		List: []*ast.Comment{c1},
	}

	cgNode := syntax.FromAstNode(cg)
	assert.NotNil(t, cgNode)
	assert.True(t, cg == cgNode.GetAstNode())
	assert.Nil(t, cgNode.GetParent())

	elmts := cgNode.GetElements()
	assert.NotNil(t, elmts)
	assert.Equal(t, 1, len(elmts))

	p := elmts[0].GetParent()
	assert.True(t, p == cgNode)
}

func TestIdentNode(t *testing.T) {
	fset := token.NewFileSet()
	fset.AddFile("test.go", 1, 1)

	n := &ast.Ident{
		NamePos: token.Pos(1),
		Name:    "Test",
	}
	node := syntax.FromAstNode(n)
	assert.NotNil(t, node)

	var bf bytes.Buffer
	ast.Fprint(&bf, fset, node, func(string, reflect.Value) bool {
		return true
	})

	fmt.Println(bf.String())
}
