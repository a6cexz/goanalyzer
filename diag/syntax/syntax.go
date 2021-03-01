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

type nodeImpl struct {
	Parent   Node
	AstNode  ast.Node
	Elements []Element
}

type tokenImpl struct {
	Pos  token.Pos
	Text string
	Kind token.Token
}
