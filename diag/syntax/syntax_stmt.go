package syntax

import (
	"go/ast"
	"go/token"
)

// BlockStmt node
type BlockStmt struct {
	*nodeImpl
	LbraceToken Token
	List        []Stmt
	RbraceToken Token
}

func (*BlockStmt) stmtNode() {}

func newBlockStmt(parent Node, node *ast.BlockStmt) *BlockStmt {
	if node == nil {
		return nil
	}
	r := &BlockStmt{}
	r.nodeImpl = getNodeImpl(parent, node)
	r.LbraceToken = newTokenByKind(r, node.Lbrace, token.LBRACE)
	r.List = newStmts(r, node.List)
	r.RbraceToken = newTokenByKind(r, node.Rbrace, token.RBRACE)
	r.Elements = getElements(r)
	return r
}

func newStmts(parent Node, nodes []ast.Stmt) []Stmt {
	if nodes == nil {
		return nil
	}
	stmts := []Stmt{}
	for _, node := range nodes {
		stmt := newStmtFromAstAndParent(parent, node)
		stmts = append(stmts, stmt)
	}
	return stmts
}

// ReturnStmt node
type ReturnStmt struct {
	*nodeImpl
	ReturnToken Token
	Results     []Expr
}

func (*ReturnStmt) stmtNode() {}

func newReturnStmt(parent Node, node *ast.ReturnStmt) *ReturnStmt {
	if node == nil {
		return nil
	}
	r := &ReturnStmt{}
	r.nodeImpl = getNodeImpl(parent, node)
	r.ReturnToken = newTokenByKind(r, node.Return, token.RETURN)
	r.Results = newExprs(r, node.Results)
	r.Elements = getElements(r)
	return r
}
