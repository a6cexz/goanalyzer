package syntaxkind

import (
	"go/ast"
)

// AstKind reprsents int kind of ast node
type AstKind int

// Ast kinds
const (
	AstNone AstKind = iota
	AstComment
	AstCommentGroup
	AstField
	AstFieldList
	AstBadExpr
	AstIdent
	AstEllipsis
	AstBasicLit
	AstFuncLit
	AstCompositeLit
	AstParenExpr
	AstSelectorExpr
	AstIndexExpr
	AstSliceExpr
	AstTypeAssertExpr
	AstCallExpr
	AstStarExpr
	AstUnaryExpr
	AstBinaryExpr
	AstKeyValueExpr
	AstArrayType
	AstStructType
	AstFuncType
	AstInterfaceType
	AstMapType
	AstChanType
	AstBadStmt
	AstDeclStmt
	AstEmptyStmt
	AstLabeledStmt
	AstExprStmt
	AstSendStmt
	AstIncDecStmt
	AstAssignStmt
	AstGoStmt
	AstDeferStmt
	AstReturnStmt
	AstBranchStmt
	AstBlockStmt
	AstIfStmt
	AstCaseClause
	AstSwitchStmt
	AstTypeSwitchStmt
	AstCommClause
	AstSelectStmt
	AstForStmt
	AstRangeStmt
	AstImportSpec
	AstValueSpec
	AstTypeSpec
	AstBadDecl
	AstGenDecl
	AstFuncDecl
	AstFile
	AstPackage
)

// GetAstKind return ast node kind
func GetAstKind(node ast.Node) AstKind {
	switch node.(type) {
	case *ast.Comment:
		return AstComment
	case *ast.CommentGroup:
		return AstCommentGroup
	case *ast.Field:
		return AstField
	case *ast.FieldList:
		return AstFieldList
	case *ast.BadExpr:
		return AstBadExpr
	case *ast.Ident:
		return AstIdent
	case *ast.Ellipsis:
		return AstEllipsis
	case *ast.BasicLit:
		return AstBasicLit
	case *ast.FuncLit:
		return AstFuncLit
	case *ast.CompositeLit:
		return AstCompositeLit
	case *ast.ParenExpr:
		return AstParenExpr
	case *ast.SelectorExpr:
		return AstSelectorExpr
	case *ast.IndexExpr:
		return AstIndexExpr
	case *ast.SliceExpr:
		return AstSliceExpr
	case *ast.TypeAssertExpr:
		return AstTypeAssertExpr
	case *ast.CallExpr:
		return AstCallExpr
	case *ast.StarExpr:
		return AstStarExpr
	case *ast.UnaryExpr:
		return AstUnaryExpr
	case *ast.BinaryExpr:
		return AstBinaryExpr
	case *ast.KeyValueExpr:
		return AstKeyValueExpr
	case *ast.ArrayType:
		return AstArrayType
	case *ast.StructType:
		return AstStructType
	case *ast.FuncType:
		return AstFuncType
	case *ast.InterfaceType:
		return AstInterfaceType
	case *ast.MapType:
		return AstMapType
	case *ast.ChanType:
		return AstChanType
	case *ast.BadStmt:
		return AstBadStmt
	case *ast.DeclStmt:
		return AstDeclStmt
	case *ast.EmptyStmt:
		return AstEmptyStmt
	case *ast.LabeledStmt:
		return AstLabeledStmt
	case *ast.ExprStmt:
		return AstExprStmt
	case *ast.SendStmt:
		return AstSendStmt
	case *ast.IncDecStmt:
		return AstIncDecStmt
	case *ast.AssignStmt:
		return AstAssignStmt
	case *ast.GoStmt:
		return AstGoStmt
	case *ast.DeferStmt:
		return AstDeferStmt
	case *ast.ReturnStmt:
		return AstReturnStmt
	case *ast.BranchStmt:
		return AstBranchStmt
	case *ast.BlockStmt:
		return AstBlockStmt
	case *ast.IfStmt:
		return AstIfStmt
	case *ast.CaseClause:
		return AstCaseClause
	case *ast.SwitchStmt:
		return AstSwitchStmt
	case *ast.TypeSwitchStmt:
		return AstTypeSwitchStmt
	case *ast.CommClause:
		return AstCommClause
	case *ast.SelectStmt:
		return AstSelectStmt
	case *ast.ForStmt:
		return AstForStmt
	case *ast.RangeStmt:
		return AstRangeStmt
	case *ast.ImportSpec:
		return AstImportSpec
	case *ast.ValueSpec:
		return AstValueSpec
	case *ast.TypeSpec:
		return AstTypeSpec
	case *ast.BadDecl:
		return AstBadDecl
	case *ast.GenDecl:
		return AstGenDecl
	case *ast.FuncDecl:
		return AstFuncDecl
	case *ast.File:
		return AstFile
	case *ast.Package:
		return AstPackage
	default:
		return AstNone
	}
}

type astVisitor func(ast.Node) bool

func (v astVisitor) Visit(node ast.Node) ast.Visitor {
	if v(node) {
		return v
	}
	return nil
}

// CollectAllNodes get all ast nodes from the root
func CollectAllNodes(root ast.Node) []ast.Node {
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

// IsExpr checks if node is expr
func IsExpr(node ast.Node) bool {
	expr := AsExpr(node)
	return expr != nil
}

// AsExpr returns ast expr node
func AsExpr(node ast.Node) ast.Expr {
	if node == nil {
		return nil
	}

	r, ok := node.(ast.Expr)
	if !ok {
		return nil
	}

	return r
}
