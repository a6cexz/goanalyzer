package syntax

import (
	"go/ast"
	"go/token"
)

func newExprs(parent Node, nodes []ast.Expr) []Expr {
	if nodes == nil {
		return nil
	}
	exprs := []Expr{}
	for _, node := range nodes {
		expr := newExprFromAstAndParent(parent, node)
		exprs = append(exprs, expr)
	}
	return exprs
}

// BadExpr node
type BadExpr struct {
	*nodeImpl
}

func (*BadExpr) exprNode() {}

func newBadExpr(parent Node, node *ast.BadExpr) *BadExpr {
	if node == nil {
		return nil
	}
	r := &BadExpr{}
	r.nodeImpl = getNodeImpl(parent, node)
	r.Elements = getElements(r)
	return r
}

// Ident node
type Ident struct {
	*nodeImpl
	NameToken Token
}

func (*Ident) exprNode() {}

func newIdent(parent Node, node *ast.Ident) *Ident {
	if node == nil {
		return nil
	}
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

// Ellipsis node
type Ellipsis struct {
	*nodeImpl
	EllipsisToken Token
	Elt           Expr
}

func (*Ellipsis) exprNode() {}

func newEllipsis(parent Node, node *ast.Ellipsis) *Ellipsis {
	if node == nil {
		return nil
	}
	r := &Ellipsis{}
	r.nodeImpl = getNodeImpl(parent, node)
	r.EllipsisToken = newTokenByKind(r, node.Ellipsis, token.ELLIPSIS)
	r.Elt = newExprFromAstAndParent(r, node.Elt)
	r.Elements = getElements(r)
	return r
}

// BasicLit node
type BasicLit struct {
	*nodeImpl
	ValueToken Token
}

func (*BasicLit) exprNode() {}

func newBasicLit(parent Node, node *ast.BasicLit) *BasicLit {
	if node == nil {
		return nil
	}
	r := &BasicLit{}
	r.nodeImpl = getNodeImpl(parent, node)
	r.ValueToken = newToken(r, node.ValuePos, node.Value, node.Kind)
	r.Elements = getElements(r)
	return r
}

// FuncLit node
type FuncLit struct {
	*nodeImpl
	Type *FuncType
	Body *BlockStmt
}

func (*FuncLit) exprNode() {}

func newFuncLit(parent Node, node *ast.FuncLit) *FuncLit {
	if node == nil {
		return nil
	}
	r := &FuncLit{}
	r.nodeImpl = getNodeImpl(parent, node)
	r.Type = newFuncType(r, node.Type)
	r.Body = newBlockStmt(r, node.Body)
	r.Elements = getElements(r)
	return r
}

// FuncType node
type FuncType struct {
	*nodeImpl
	FuncToken Token
	Params    *FieldList
	Results   *FieldList
}

func (*FuncType) exprNode() {}

func newFuncType(parent Node, node *ast.FuncType) *FuncType {
	if node == nil {
		return nil
	}
	r := &FuncType{}
	r.nodeImpl = getNodeImpl(parent, node)
	r.FuncToken = newTokenByKind(r, node.Func, token.FUNC)
	r.Params = newFieldList(r, node.Params)
	r.Results = newFieldList(r, node.Results)
	r.Elements = getElements(r)
	return r
}
