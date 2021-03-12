package syntax_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"maxmapps.com/goanalyzer/diag/syntax"
)

func TestNewCommentGroupNode(t *testing.T) {
	commentGroup := getCommentGroup("Test1", "Test2")

	elmt := syntax.NewElementFromAst(commentGroup)
	assert.NotNil(t, elmt)
	assert.True(t, commentGroup == elmt.GetAstNode())
	assert.Nil(t, elmt.GetParent())

	elmts := elmt.GetElements()
	assert.NotNil(t, elmts)
	assert.Equal(t, 2, len(elmts))

	parent := elmts[0].GetParent()
	assert.True(t, parent == elmt)

	e := `node *ast.CommentGroup
parent <nil>
elmnts: [
	node *ast.Comment
	parent *ast.CommentGroup
	elmnts: []

	node *ast.Comment
	parent *ast.CommentGroup
	elmnts: []
]
`
	checkSyntaxTree2(t, e, commentGroup)
}
