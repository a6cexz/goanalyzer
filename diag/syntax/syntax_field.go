package syntax

import (
	"go/ast"
	"go/token"
)

// Field node
type Field struct {
	*nodeImpl
	Doc     *CommentGroup
	Names   []*Ident
	Type    Expr
	Tag     *BasicLit
	Comment *CommentGroup
}

func newField(parent Node, node *ast.Field) *Field {
	if node == nil {
		return nil
	}
	r := &Field{}
	r.nodeImpl = getNodeImpl(parent, node)
	r.Doc = newCommentGroup(r, node.Doc)
	r.Names = newIdents(r, node.Names)
	r.Type = newExprFromAstAndParent(r, node.Type)
	r.Tag = newBasicLit(r, node.Tag)
	r.Comment = newCommentGroup(r, node.Comment)
	r.Elements = getElements(r)
	return r
}

func newFields(parent Node, nodes []*ast.Field) []*Field {
	if nodes == nil {
		return nil
	}
	fields := []*Field{}
	for _, node := range nodes {
		field := newField(parent, node)
		fields = append(fields, field)
	}
	return fields
}

// FieldList node
type FieldList struct {
	*nodeImpl
	Opening Token
	List    []*Field
	Closing Token
}

func newFieldList(parent Node, node *ast.FieldList) *FieldList {
	if node == nil {
		return nil
	}
	r := &FieldList{}
	r.nodeImpl = getNodeImpl(parent, node)
	r.Opening = newTokenByKind(r, node.Opening, token.LPAREN)
	r.List = newFields(r, node.List)
	r.Closing = newTokenByKind(r, node.Closing, token.RPAREN)
	r.Elements = getElements(r)
	return r
}
