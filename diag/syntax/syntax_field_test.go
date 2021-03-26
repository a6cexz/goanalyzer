package syntax_test

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestFieldNode(t *testing.T) {
	e := `node *ast.Field
parent <nil>
elmnts: [
	node *ast.CommentGroup
	parent *ast.Field
	elmnts: [
		node *ast.Comment
		parent *ast.CommentGroup
		elmnts: []
	]

	node *ast.Ident
	parent *ast.Field
	elmnts: [
		token a IDENT
	]

	node *ast.Ident
	parent *ast.Field
	elmnts: [
		token b IDENT
	]

	node *ast.Ident
	parent *ast.Field
	elmnts: [
		token int IDENT
	]

	node *ast.BasicLit
	parent *ast.Field
	elmnts: [
		token lit STRING
	]

	node *ast.CommentGroup
	parent *ast.Field
	elmnts: [
		node *ast.Comment
		parent *ast.CommentGroup
		elmnts: []
	]
]
`

	doc := getCommentGroup("test")
	f := &ast.Field{
		Doc:     doc,
		Names:   getIdents("a", "b"),
		Type:    getIdent("int"),
		Tag:     getBasicLit(token.STRING, "lit"),
		Comment: getCommentGroup("line"),
	}
	checkSyntaxTree2(t, e, f)
}

func TestFieldListNode(t *testing.T) {
	e := `node *ast.FieldList
parent <nil>
elmnts: [
	token ( (

	node *ast.Field
	parent *ast.FieldList
	elmnts: [
		node *ast.Ident
		parent *ast.Field
		elmnts: [
			token a IDENT
		]
	
		node *ast.Ident
		parent *ast.Field
		elmnts: [
			token int IDENT
		]
	]

	token ) )
]
`

	f := &ast.Field{
		Names: getIdents("a"),
		Type:  getIdent("int"),
	}
	fieldList := &ast.FieldList{
		Opening: token.Pos(1),
		List:    []*ast.Field{f},
		Closing: token.Pos(2),
	}
	checkSyntaxTree2(t, e, fieldList)
}
