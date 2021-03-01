package syntaxkind_test

import (
	"fmt"
	"go/ast"
	"testing"

	"github.com/a6cexz/goanalyzer/diag/syntax/asttest"
	"github.com/a6cexz/goanalyzer/diag/syntax/syntaxkind"
	"github.com/stretchr/testify/assert"
)

var (
	TypeAstKindMap = map[string]syntaxkind.AstKind{
		"*ast.Comment":        syntaxkind.AstComment,
		"*ast.Field":          syntaxkind.AstField,
		"*ast.FieldList":      syntaxkind.AstFieldList,
		"*ast.BadExpr":        syntaxkind.AstBadExpr,
		"*ast.Ident":          syntaxkind.AstIdent,
		"*ast.Ellipsis":       syntaxkind.AstEllipsis,
		"*ast.BasicLit":       syntaxkind.AstBasicLit,
		"*ast.FuncLit":        syntaxkind.AstFuncLit,
		"*ast.CompositeLit":   syntaxkind.AstCompositeLit,
		"*ast.ParenExpr":      syntaxkind.AstParenExpr,
		"*ast.SelectorExpr":   syntaxkind.AstSelectorExpr,
		"*ast.IndexExpr":      syntaxkind.AstIndexExpr,
		"*ast.SliceExpr":      syntaxkind.AstSliceExpr,
		"*ast.TypeAssertExpr": syntaxkind.AstTypeAssertExpr,
		"*ast.CallExpr":       syntaxkind.AstCallExpr,
		"*ast.StarExpr":       syntaxkind.AstStarExpr,
		"*ast.UnaryExpr":      syntaxkind.AstUnaryExpr,
		"*ast.BinaryExpr":     syntaxkind.AstBinaryExpr,
		"*ast.KeyValueExpr":   syntaxkind.AstKeyValueExpr,
		"*ast.ArrayType":      syntaxkind.AstArrayType,
		"*ast.StructType":     syntaxkind.AstStructType,
		"*ast.FuncType":       syntaxkind.AstFuncType,
		"*ast.InterfaceType":  syntaxkind.AstInterfaceType,
		"*ast.MapType":        syntaxkind.AstMapType,
		"*ast.ChanType":       syntaxkind.AstChanType,
		"*ast.BadStmt":        syntaxkind.AstBadStmt,
		"*ast.DeclStmt":       syntaxkind.AstDeclStmt,
		"*ast.EmptyStmt":      syntaxkind.AstEmptyStmt,
		"*ast.LabeledStmt":    syntaxkind.AstLabeledStmt,
		"*ast.ExprStmt":       syntaxkind.AstExprStmt,
		"*ast.SendStmt":       syntaxkind.AstSendStmt,
		"*ast.IncDecStmt":     syntaxkind.AstIncDecStmt,
		"*ast.AssignStmt":     syntaxkind.AstAssignStmt,
		"*ast.GoStmt":         syntaxkind.AstGoStmt,
		"*ast.DeferStmt":      syntaxkind.AstDeferStmt,
		"*ast.ReturnStmt":     syntaxkind.AstReturnStmt,
		"*ast.BranchStmt":     syntaxkind.AstBranchStmt,
		"*ast.BlockStmt":      syntaxkind.AstBlockStmt,
		"*ast.IfStmt":         syntaxkind.AstIfStmt,
		"*ast.CaseClause":     syntaxkind.AstCaseClause,
		"*ast.SwitchStmt":     syntaxkind.AstSwitchStmt,
		"*ast.TypeSwitchStmt": syntaxkind.AstTypeSwitchStmt,
		"*ast.CommClause":     syntaxkind.AstCommClause,
		"*ast.SelectStmt":     syntaxkind.AstSelectStmt,
		"*ast.ForStmt":        syntaxkind.AstForStmt,
		"*ast.RangeStmt":      syntaxkind.AstRangeStmt,
		"*ast.ImportSpec":     syntaxkind.AstImportSpec,
		"*ast.ValueSpec":      syntaxkind.AstValueSpec,
		"*ast.TypeSpec":       syntaxkind.AstTypeSpec,
		"*ast.BadDecl":        syntaxkind.AstBadDecl,
		"*ast.GenDecl":        syntaxkind.AstGenDecl,
		"*ast.FuncDecl":       syntaxkind.AstFuncDecl,
		"*ast.File":           syntaxkind.AstFile,
		"*ast.Package":        syntaxkind.AstPackage,
	}
)

func checkNodeAstKinds(t *testing.T, src string) {
	_, file, _, err := asttest.ParseTestFile(src)
	assert.NoError(t, err)

	nodes := syntaxkind.CollectAllNodes(file)
	for _, node := range nodes {
		checkNodeAstKind(t, node)
	}
}
func checkNodeAstKind(t *testing.T, node ast.Node) {
	nodeType := fmt.Sprintf("%T", node)
	if testKind, ok := TypeAstKindMap[nodeType]; ok {
		assert.Equal(t, testKind, syntaxkind.GetAstKind(node))
	} else {
		assert.Failf(t, "Error", "Found node type without ast kind: %v", nodeType)
	}
}

func TestNodeAstKinds(t *testing.T) {
	src := `package main

import (
	"fmt"
)

// Comment1
// Comment2

const SomeConst = 10
var SomeVar = 10
type SomeStruct struct {
	Field1 string
	Field2 string
}
type SomeInterface interface {
	Func1() int
	Func2() float64
}
func (s *SomeStruct) Func1() int {
	if err != nil {
		return err
	}

	for i, item := range items {
		if i == 0 || i < 0 {
			break
		} else {
			continue
		}
	}

	return 10
}`
	checkNodeAstKinds(t, src)
}
