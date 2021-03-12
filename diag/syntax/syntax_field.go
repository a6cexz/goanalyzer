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

// Ident node
type Ident struct {
	*nodeImpl
	NameToken Token
}

func (*Ident) exprNode() {}

func newToken(parent Node, pos token.Pos, text string, kind token.Token) Token {
	r := &tokenImpl{}
	r.Parent = parent
	r.Pos = pos
	r.Text = text
	r.Kind = kind
	return r
}

func newIdent(parent Node, node *ast.Ident) *Ident {
	r := &Ident{}
	r.nodeImpl = getNodeImpl(parent, node)
	r.NameToken = newToken(r, node.NamePos, node.Name, token.IDENT)
	r.Elements = getElements(r)
	return r
}

func newIdents(parent Node, nodes []*ast.Ident) []*Ident {
	if nodes == nil {
		return nil
	}
	idents := []*Ident{}
	for _, node := range nodes {
		ident := newIdent(parent, node)
		idents = append(idents, ident)
	}
	return idents
}

// BasicLit node
type BasicLit struct {
	*nodeImpl
	ValueToken Token
}

func (*BasicLit) exprNode() {}

func newBasicLit(parent Node, node *ast.BasicLit) *BasicLit {
	r := &BasicLit{}
	r.nodeImpl = getNodeImpl(parent, node)
	r.ValueToken = newToken(r, node.ValuePos, node.Value, node.Kind)
	r.Elements = getElements(r)
	return r
}
