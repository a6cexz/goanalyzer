package syntax

import (
	"go/ast"
	"go/token"
)

// Element represents node or token
type Element interface {
	GetParent() Node
}

// Node represents syntax node
type Node interface {
	Element
	GetAstNode() ast.Node
	GetElements() []Element
}

// Token represents token node
type Token interface {
	Element
}

// FromAstNode returns node from the given ast.Node
func FromAstNode(node ast.Node) Node {
	return fromAstNodeAndParent(nil, node)
}

type nodeImpl struct {
	Parent   Node
	AstNode  ast.Node
	Elements []Element
}

func (n *nodeImpl) GetParent() Node {
	return n.Parent
}

func (n *nodeImpl) GetAstNode() ast.Node {
	return n.AstNode
}

func (n *nodeImpl) GetElements() []Element {
	return n.Elements
}

type tokenImpl struct {
	Parent Node
	Pos    token.Pos
	Text   string
	Kind   token.Token
}

func (t *tokenImpl) GetParent() Node {
	return t.Parent
}

func fromAstNodeAndParent(parent Node, node ast.Node) Node {
	if node == nil {
		return nil
	}
	n := &nodeImpl{
		Parent:  parent,
		AstNode: node,
	}
	n.Elements = loadElements(n, node)
	return n
}

func tokenNode(parent Node, pos token.Pos, text string, kind token.Token) Token {
	t := &tokenImpl{
		Parent: parent,
		Pos:    pos,
		Text:   text,
		Kind:   kind,
	}
	return t
}

func loadElements(parent Node, node ast.Node) []Element {
	switch n := node.(type) {
	case *ast.Comment:
		return nil

	case *ast.CommentGroup:
		elmts := []Element{}
		for _, c := range n.List {
			cn := fromAstNodeAndParent(parent, c)
			elmts = append(elmts, cn)
		}
		return elmts

	case *ast.Ident:
		elmts := []Element{}
		tn := tokenNode(parent, n.NamePos, n.Name, token.IDENT)
		elmts = append(elmts, tn)
		return elmts
	}
	return nil
}
