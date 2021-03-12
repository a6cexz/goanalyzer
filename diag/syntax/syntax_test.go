package syntax_test

import (
	"bytes"
	"go/ast"
	"go/token"
	"testing"

	"github.com/stretchr/testify/assert"
	"maxmapps.com/goanalyzer/diag/syntax"
)

func printSyntaxTree(node ast.Node) {
	elmt := syntax.FromAstNode(node)
	syntax.Print(elmt)
}
func checkSyntaxTree(t *testing.T, expected string, node ast.Node) {
	elmt := syntax.FromAstNode(node)
	var buffer bytes.Buffer
	syntax.PrintTo(&buffer, elmt)
	str := buffer.String()
	assert.Equal(t, expected, str)
}

func printSyntaxTree2(node ast.Node) {
	elmt := syntax.NewElementFromAst(node)
	syntax.Print(elmt)
}
func checkSyntaxTree2(t *testing.T, expected string, node ast.Node) {
	elmt := syntax.NewElementFromAst(node)
	var buffer bytes.Buffer
	syntax.PrintTo(&buffer, elmt)
	str := buffer.String()
	assert.Equal(t, expected, str)
}

func getComment(text string) *ast.Comment {
	c := &ast.Comment{
		Slash: token.Pos(1),
		Text:  text,
	}
	return c
}
func getComments(texts ...string) []*ast.Comment {
	r := []*ast.Comment{}
	for _, text := range texts {
		c := getComment(text)
		r = append(r, c)
	}
	return r
}
func getCommentGroup(texts ...string) *ast.CommentGroup {
	list := getComments(texts...)
	r := &ast.CommentGroup{
		List: list,
	}
	return r
}

func getField(name string, fieldType string) *ast.Field {
	r := &ast.Field{
		Names: getIdents(name),
		Type:  getIdent(fieldType),
	}
	return r
}

func getFieldList(fields ...*ast.Field) *ast.FieldList {
	r := &ast.FieldList{
		Opening: token.Pos(1),
		List:    fields,
		Closing: token.Pos(2),
	}
	return r
}

func getBadExpr() *ast.BadExpr {
	r := &ast.BadExpr{
		From: token.Pos(1),
		To:   token.Pos(2),
	}
	return r
}

func getIdent(name string) *ast.Ident {
	r := &ast.Ident{
		NamePos: token.Pos(1),
		Name:    name,
	}
	return r
}
func getIdents(names ...string) []*ast.Ident {
	r := []*ast.Ident{}
	for _, text := range names {
		ident := getIdent(text)
		r = append(r, ident)
	}
	return r
}

func getEllipsis(name string) *ast.Ellipsis {
	r := &ast.Ellipsis{
		Ellipsis: token.Pos(1),
		Elt:      getIdent(name),
	}
	return r
}

func getBasicLit(kind token.Token, value string) *ast.BasicLit {
	r := &ast.BasicLit{
		ValuePos: token.Pos(1),
		Kind:     kind,
		Value:    value,
	}
	return r
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
	checkSyntaxTree(t, e, fieldList)
}

func TestBadExprNode(t *testing.T) {
	e := `node *ast.BadExpr
parent <nil>
elmnts: []
`
	expr := getBadExpr()
	checkSyntaxTree(t, e, expr)
}

func TestIdentNode(t *testing.T) {
	e := `node *ast.Ident
parent <nil>
elmnts: [
	token test IDENT
]
`
	n := getIdent("test")
	checkSyntaxTree(t, e, n)
}

func TestEllipsisNode(t *testing.T) {
	e := `node *ast.Ellipsis
parent <nil>
elmnts: [
	token ... ...

	node *ast.Ident
	parent *ast.Ellipsis
	elmnts: [
		token test IDENT
	]
]
`
	n := getEllipsis("test")
	checkSyntaxTree(t, e, n)
}

func TestBasicLitNode(t *testing.T) {
	e := `node *ast.BasicLit
parent <nil>
elmnts: [
	token test STRING
]
`
	n := getBasicLit(token.STRING, "test")
	checkSyntaxTree(t, e, n)
}

func TestFuncLitNode(t *testing.T) {
	e := `node *ast.FuncLit
parent <nil>
elmnts: [
	node *ast.FuncType
	parent *ast.FuncLit
	elmnts: [
		token func func
	
		node *ast.FieldList
		parent *ast.FuncType
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
	
		node *ast.FieldList
		parent *ast.FuncType
		elmnts: [
			token ( (
		
			node *ast.Field
			parent *ast.FieldList
			elmnts: [
				node *ast.Ident
				parent *ast.Field
				elmnts: [
					token r IDENT
				]
			
				node *ast.Ident
				parent *ast.Field
				elmnts: [
					token string IDENT
				]
			]
		
			token ) )
		]
	]

	node *ast.BlockStmt
	parent *ast.FuncLit
	elmnts: [
		token { {
	
		node *ast.ReturnStmt
		parent *ast.BlockStmt
		elmnts: [
			token return return
		
			node *ast.Ident
			parent *ast.ReturnStmt
			elmnts: [
				token a IDENT
			]
		]
	
		token } }
	]
]
`

	ft := &ast.FuncType{
		Func:    token.Pos(1),
		Params:  getFieldList(getField("a", "int")),
		Results: getFieldList(getField("r", "string")),
	}

	ret := &ast.ReturnStmt{
		Return: token.Pos(1),
		Results: []ast.Expr{
			getIdent("a"),
		},
	}

	b := &ast.BlockStmt{
		Lbrace: token.Pos(1),
		List:   []ast.Stmt{ret},
		Rbrace: token.Pos(2),
	}

	n := &ast.FuncLit{
		Type: ft,
		Body: b,
	}

	checkSyntaxTree(t, e, n)
}

func TestCompositeLitNode(t *testing.T) {
	e := `node *ast.CompositeLit
parent <nil>
elmnts: [
	node *ast.Ident
	parent *ast.CompositeLit
	elmnts: [
		token test IDENT
	]

	token { {

	node *ast.Ident
	parent *ast.CompositeLit
	elmnts: [
		token a IDENT
	]

	token } }
]
`
	n := &ast.CompositeLit{
		Type:   getIdent("test"),
		Lbrace: token.Pos(1),
		Elts:   []ast.Expr{getIdent("a")},
		Rbrace: token.Pos(2),
	}
	checkSyntaxTree(t, e, n)
}

func TestParenExprNode(t *testing.T) {
	e := `node *ast.ParenExpr
parent <nil>
elmnts: [
	token ( (

	node *ast.Ident
	parent *ast.ParenExpr
	elmnts: [
		token name IDENT
	]

	token ) )
]
`
	n := &ast.ParenExpr{
		Lparen: token.Pos(1),
		X:      getIdent("name"),
		Rparen: token.Pos(2),
	}
	checkSyntaxTree(t, e, n)
}

func TestSelectorExprNode(t *testing.T) {
	e := `node *ast.SelectorExpr
parent <nil>
elmnts: [
	node *ast.Ident
	parent *ast.SelectorExpr
	elmnts: [
		token x IDENT
	]

	node *ast.Ident
	parent *ast.SelectorExpr
	elmnts: [
		token sel IDENT
	]
]
`
	n := &ast.SelectorExpr{
		X:   getIdent("x"),
		Sel: getIdent("sel"),
	}
	checkSyntaxTree(t, e, n)
}

func TestIndexExprNode(t *testing.T) {
	e := `node *ast.IndexExpr
parent <nil>
elmnts: [
	node *ast.Ident
	parent *ast.IndexExpr
	elmnts: [
		token x IDENT
	]

	token [ [

	node *ast.Ident
	parent *ast.IndexExpr
	elmnts: [
		token index IDENT
	]

	token ] ]
]
`
	n := &ast.IndexExpr{
		X:      getIdent("x"),
		Lbrack: token.Pos(1),
		Index:  getIdent("index"),
		Rbrack: token.Pos(2),
	}
	checkSyntaxTree(t, e, n)
}

func TestSliceExprNode(t *testing.T) {
	e := `node *ast.SliceExpr
parent <nil>
elmnts: [
	node *ast.Ident
	parent *ast.SliceExpr
	elmnts: [
		token x IDENT
	]

	token [ [

	node *ast.Ident
	parent *ast.SliceExpr
	elmnts: [
		token low IDENT
	]

	node *ast.Ident
	parent *ast.SliceExpr
	elmnts: [
		token high IDENT
	]

	node *ast.Ident
	parent *ast.SliceExpr
	elmnts: [
		token max IDENT
	]

	token ] ]
]
`
	n := &ast.SliceExpr{
		X:      getIdent("x"),
		Lbrack: token.Pos(1),
		Low:    getIdent("low"),
		High:   getIdent("high"),
		Max:    getIdent("max"),
		Rbrack: token.Pos(2),
	}
	checkSyntaxTree(t, e, n)
}

func TestTypeAssertExprNode(t *testing.T) {
	e := `node *ast.TypeAssertExpr
parent <nil>
elmnts: [
	node *ast.Ident
	parent *ast.TypeAssertExpr
	elmnts: [
		token x IDENT
	]

	token ( (

	node *ast.Ident
	parent *ast.TypeAssertExpr
	elmnts: [
		token type IDENT
	]

	token ) )
]
`
	n := &ast.TypeAssertExpr{
		X:      getIdent("x"),
		Lparen: token.Pos(1),
		Type:   getIdent("type"),
		Rparen: token.Pos(2),
	}
	checkSyntaxTree(t, e, n)
}

func TestCallExprNode(t *testing.T) {
	e := `node *ast.CallExpr
parent <nil>
elmnts: [
	node *ast.Ident
	parent *ast.CallExpr
	elmnts: [
		token fun IDENT
	]

	token ( (

	node *ast.Ident
	parent *ast.CallExpr
	elmnts: [
		token arg1 IDENT
	]

	token ... ...

	token ) )
]
`
	n := &ast.CallExpr{
		Fun:      getIdent("fun"),
		Lparen:   token.Pos(1),
		Args:     []ast.Expr{getIdent("arg1")},
		Ellipsis: token.Pos(2),
		Rparen:   token.Pos(3),
	}
	checkSyntaxTree(t, e, n)
}

func TestStarExprNode(t *testing.T) {
	e := `node *ast.StarExpr
parent <nil>
elmnts: [
	token * *

	node *ast.Ident
	parent *ast.StarExpr
	elmnts: [
		token arg1 IDENT
	]
]
`
	n := &ast.StarExpr{
		Star: token.Pos(1),
		X:    getIdent("arg1"),
	}
	checkSyntaxTree(t, e, n)
}

func TestUnaryExprNode(t *testing.T) {
	e := `node *ast.UnaryExpr
parent <nil>
elmnts: [
	token * *

	node *ast.Ident
	parent *ast.UnaryExpr
	elmnts: [
		token x IDENT
	]
]
`
	n := &ast.UnaryExpr{
		OpPos: token.Pos(1),
		Op:    token.MUL,
		X:     getIdent("x"),
	}
	checkSyntaxTree(t, e, n)
}

func TestBinaryExprNode(t *testing.T) {
	e := `node *ast.BinaryExpr
parent <nil>
elmnts: [
	node *ast.Ident
	parent *ast.BinaryExpr
	elmnts: [
		token x IDENT
	]

	token * *

	node *ast.Ident
	parent *ast.BinaryExpr
	elmnts: [
		token y IDENT
	]
]
`
	n := &ast.BinaryExpr{
		X:     getIdent("x"),
		OpPos: token.Pos(1),
		Op:    token.MUL,
		Y:     getIdent("y"),
	}
	checkSyntaxTree(t, e, n)
}

func TestKeyValueExprNode(t *testing.T) {
	e := `node *ast.KeyValueExpr
parent <nil>
elmnts: [
	node *ast.Ident
	parent *ast.KeyValueExpr
	elmnts: [
		token key IDENT
	]

	token : :

	node *ast.Ident
	parent *ast.KeyValueExpr
	elmnts: [
		token value IDENT
	]
]
`
	n := &ast.KeyValueExpr{
		Key:   getIdent("key"),
		Colon: token.Pos(1),
		Value: getIdent("value"),
	}
	checkSyntaxTree(t, e, n)
}

func TestArrayTypeNode(t *testing.T) {
	e := `node *ast.ArrayType
parent <nil>
elmnts: [
	token [ [

	node *ast.BasicLit
	parent *ast.ArrayType
	elmnts: [
		token 1 INT
	]

	node *ast.Ident
	parent *ast.ArrayType
	elmnts: [
		token test IDENT
	]
]
`
	n := &ast.ArrayType{
		Lbrack: token.Pos(1),
		Len:    getBasicLit(token.INT, "1"),
		Elt:    getIdent("test"),
	}
	checkSyntaxTree(t, e, n)
}

func TestStructTypeNode(t *testing.T) {
	e := `node *ast.StructType
parent <nil>
elmnts: [
	token struct struct

	node *ast.FieldList
	parent *ast.StructType
	elmnts: [
		token ( (
	
		node *ast.Field
		parent *ast.FieldList
		elmnts: [
			node *ast.Ident
			parent *ast.Field
			elmnts: [
				token name IDENT
			]
		
			node *ast.Ident
			parent *ast.Field
			elmnts: [
				token string IDENT
			]
		]
	
		token ) )
	]
]
`
	n := &ast.StructType{
		Struct: token.Pos(1),
		Fields: getFieldList(getField("name", "string")),
	}
	checkSyntaxTree(t, e, n)
}

func TestFuncTypeNode(t *testing.T) {
	e := `node *ast.FuncType
parent <nil>
elmnts: [
	token func func

	node *ast.FieldList
	parent *ast.FuncType
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

	node *ast.FieldList
	parent *ast.FuncType
	elmnts: [
		token ( (
	
		node *ast.Field
		parent *ast.FieldList
		elmnts: [
			node *ast.Ident
			parent *ast.Field
			elmnts: [
				token r IDENT
			]
		
			node *ast.Ident
			parent *ast.Field
			elmnts: [
				token string IDENT
			]
		]
	
		token ) )
	]
]
`
	n := &ast.FuncType{
		Func:    token.Pos(1),
		Params:  getFieldList(getField("a", "int")),
		Results: getFieldList(getField("r", "string")),
	}
	checkSyntaxTree(t, e, n)
}

func TestInterfaceTypeNode(t *testing.T) {
	e := `node *ast.InterfaceType
parent <nil>
elmnts: [
	token interface interface

	node *ast.FieldList
	parent *ast.InterfaceType
	elmnts: [
		token ( (
	
		node *ast.Field
		parent *ast.FieldList
		elmnts: [
			node *ast.Ident
			parent *ast.Field
			elmnts: [
				token m1 IDENT
			]
		
			node *ast.Ident
			parent *ast.Field
			elmnts: [
				token int IDENT
			]
		]
	
		token ) )
	]
]
`
	n := &ast.InterfaceType{
		Interface: token.Pos(1),
		Methods:   getFieldList(getField("m1", "int")),
	}
	checkSyntaxTree(t, e, n)
}

func TestMapTypeNode(t *testing.T) {
	e := `node *ast.MapType
parent <nil>
elmnts: [
	token map map

	node *ast.Ident
	parent *ast.MapType
	elmnts: [
		token key IDENT
	]

	node *ast.Ident
	parent *ast.MapType
	elmnts: [
		token value IDENT
	]
]
`
	n := &ast.MapType{
		Map:   token.Pos(1),
		Key:   getIdent("key"),
		Value: getIdent("value"),
	}
	checkSyntaxTree(t, e, n)
}

func TestChanTypeNode1(t *testing.T) {
	e := `node *ast.ChanType
parent <nil>
elmnts: [
	token chan chan

	token <- <-

	node *ast.Ident
	parent *ast.ChanType
	elmnts: [
		token value IDENT
	]
]
`
	n := &ast.ChanType{
		Begin: token.Pos(1),
		Arrow: token.Pos(2),
		Dir:   ast.SEND,
		Value: getIdent("value"),
	}
	checkSyntaxTree(t, e, n)
}

func TestChanTypeNode2(t *testing.T) {
	e := `node *ast.ChanType
parent <nil>
elmnts: [
	token <- <-

	node *ast.Ident
	parent *ast.ChanType
	elmnts: [
		token value IDENT
	]
]
`
	n := &ast.ChanType{
		Begin: token.Pos(1),
		Arrow: token.Pos(1),
		Dir:   ast.SEND,
		Value: getIdent("value"),
	}
	checkSyntaxTree(t, e, n)
}

func TestBadStmtNode(t *testing.T) {
	e := `node *ast.BadStmt
parent <nil>
elmnts: []
`
	n := &ast.BadStmt{
		From: token.Pos(1),
		To:   token.Pos(2),
	}
	checkSyntaxTree(t, e, n)
}

func TestDeclStmtNode(t *testing.T) {
	e := `node *ast.DeclStmt
parent <nil>
elmnts: [
	node *ast.BadDecl
	parent *ast.DeclStmt
	elmnts: []
]
`
	n := &ast.DeclStmt{
		Decl: &ast.BadDecl{
			From: token.Pos(1),
			To:   token.Pos(2),
		},
	}
	checkSyntaxTree(t, e, n)
}

func TestEmptyStmtNode1(t *testing.T) {
	e := `node *ast.EmptyStmt
parent <nil>
elmnts: [
	token ; ;
]
`
	n := &ast.EmptyStmt{
		Semicolon: token.Pos(1),
		Implicit:  false,
	}
	checkSyntaxTree(t, e, n)
}

func TestEmptyStmtNode2(t *testing.T) {
	e := `node *ast.EmptyStmt
parent <nil>
elmnts: []
`
	n := &ast.EmptyStmt{
		Semicolon: token.Pos(1),
		Implicit:  true,
	}
	checkSyntaxTree(t, e, n)
}

func TestLabeledStmtNode(t *testing.T) {
	e := `node *ast.LabeledStmt
parent <nil>
elmnts: [
	node *ast.Ident
	parent *ast.LabeledStmt
	elmnts: [
		token label IDENT
	]

	token : :

	node *ast.BadStmt
	parent *ast.LabeledStmt
	elmnts: []
]
`
	n := &ast.LabeledStmt{
		Label: getIdent("label"),
		Colon: token.Pos(1),
		Stmt: &ast.BadStmt{
			From: token.Pos(2),
			To:   token.Pos(3),
		},
	}
	checkSyntaxTree(t, e, n)
}

func TestExprStmtNode(t *testing.T) {
	e := `node *ast.ExprStmt
parent <nil>
elmnts: [
	node *ast.Ident
	parent *ast.ExprStmt
	elmnts: [
		token x IDENT
	]
]
`
	n := &ast.ExprStmt{
		X: getIdent("x"),
	}
	checkSyntaxTree(t, e, n)
}

func TestSendStmtNode(t *testing.T) {
	e := `node *ast.SendStmt
parent <nil>
elmnts: [
	node *ast.Ident
	parent *ast.SendStmt
	elmnts: [
		token ch IDENT
	]

	token <- <-

	node *ast.Ident
	parent *ast.SendStmt
	elmnts: [
		token value IDENT
	]
]
`
	n := &ast.SendStmt{
		Chan:  getIdent("ch"),
		Arrow: token.Pos(1),
		Value: getIdent("value"),
	}
	checkSyntaxTree(t, e, n)
}

func TestAssignStmtNode(t *testing.T) {
	e := `node *ast.AssignStmt
parent <nil>
elmnts: [
	node *ast.Ident
	parent *ast.AssignStmt
	elmnts: [
		token left IDENT
	]

	token = =

	node *ast.Ident
	parent *ast.AssignStmt
	elmnts: [
		token right IDENT
	]
]
`
	n := &ast.AssignStmt{
		Lhs:    []ast.Expr{getIdent("left")},
		TokPos: token.Pos(1),
		Tok:    token.ASSIGN,
		Rhs:    []ast.Expr{getIdent("right")},
	}
	checkSyntaxTree(t, e, n)
}

func TestGoStmtNode(t *testing.T) {
	e := `node *ast.GoStmt
parent <nil>
elmnts: [
	token go go

	node *ast.CallExpr
	parent *ast.GoStmt
	elmnts: [
		node *ast.Ident
		parent *ast.CallExpr
		elmnts: [
			token f IDENT
		]
	
		token ( (
	
		node *ast.Ident
		parent *ast.CallExpr
		elmnts: [
			token a IDENT
		]
	
		token ... ...
	
		token ) )
	]
]
`
	n := &ast.GoStmt{
		Go: token.Pos(1),
		Call: &ast.CallExpr{
			Fun:      getIdent("f"),
			Lparen:   token.Pos(2),
			Args:     []ast.Expr{getIdent("a")},
			Ellipsis: token.Pos(3),
			Rparen:   token.Pos(4),
		},
	}
	checkSyntaxTree(t, e, n)
}

func TestDeferStmtNode(t *testing.T) {
	e := `node *ast.DeferStmt
parent <nil>
elmnts: [
	token defer defer

	node *ast.CallExpr
	parent *ast.DeferStmt
	elmnts: [
		node *ast.Ident
		parent *ast.CallExpr
		elmnts: [
			token f IDENT
		]
	
		token ( (
	
		node *ast.Ident
		parent *ast.CallExpr
		elmnts: [
			token a IDENT
		]
	
		token ... ...
	
		token ) )
	]
]
`
	n := &ast.DeferStmt{
		Defer: token.Pos(1),
		Call: &ast.CallExpr{
			Fun:      getIdent("f"),
			Lparen:   token.Pos(2),
			Args:     []ast.Expr{getIdent("a")},
			Ellipsis: token.Pos(3),
			Rparen:   token.Pos(4),
		},
	}
	checkSyntaxTree(t, e, n)
}

func TestReturnStmtNode(t *testing.T) {
	e := `node *ast.ReturnStmt
parent <nil>
elmnts: [
	token return return

	node *ast.Ident
	parent *ast.ReturnStmt
	elmnts: [
		token r IDENT
	]
]
`
	n := &ast.ReturnStmt{
		Return:  token.Pos(1),
		Results: []ast.Expr{getIdent("r")},
	}
	checkSyntaxTree(t, e, n)
}

func TestBranchStmtNode(t *testing.T) {
	e := `node *ast.BranchStmt
parent <nil>
elmnts: [
	token break break

	node *ast.Ident
	parent *ast.BranchStmt
	elmnts: [
		token label IDENT
	]
]
`
	n := &ast.BranchStmt{
		TokPos: token.Pos(1),
		Tok:    token.BREAK,
		Label:  getIdent("label"),
	}
	checkSyntaxTree(t, e, n)
}

func TestBlockStmtNode(t *testing.T) {
	e := `node *ast.BlockStmt
parent <nil>
elmnts: [
	token { {

	node *ast.BadStmt
	parent *ast.BlockStmt
	elmnts: []

	token } }
]
`
	n := &ast.BlockStmt{
		Lbrace: token.Pos(1),
		List:   []ast.Stmt{&ast.BadStmt{}},
		Rbrace: token.Pos(2),
	}
	checkSyntaxTree(t, e, n)
}

func TestIfStmtNode(t *testing.T) {
	e := `node *ast.IfStmt
parent <nil>
elmnts: [
	token if if

	node *ast.BadStmt
	parent *ast.IfStmt
	elmnts: []

	node *ast.Ident
	parent *ast.IfStmt
	elmnts: [
		token cond IDENT
	]

	node *ast.BlockStmt
	parent *ast.IfStmt
	elmnts: [
		token { {
	
		token } }
	]

	node *ast.BadStmt
	parent *ast.IfStmt
	elmnts: []
]
`
	n := &ast.IfStmt{
		If:   token.Pos(1),
		Init: &ast.BadStmt{},
		Cond: getIdent("cond"),
		Body: &ast.BlockStmt{},
		Else: &ast.BadStmt{},
	}
	checkSyntaxTree(t, e, n)
}

func TestCaseClauseNode(t *testing.T) {
	e := `node *ast.CaseClause
parent <nil>
elmnts: [
	token case case

	node *ast.Ident
	parent *ast.CaseClause
	elmnts: [
		token cond IDENT
	]

	token : :

	node *ast.BadStmt
	parent *ast.CaseClause
	elmnts: []
]
`
	n := &ast.CaseClause{
		Case:  token.Pos(1),
		List:  []ast.Expr{getIdent("cond")},
		Colon: token.Pos(2),
		Body:  []ast.Stmt{&ast.BadStmt{}},
	}
	checkSyntaxTree(t, e, n)
}

func TestSwitchStmtNode(t *testing.T) {
	e := `node *ast.SwitchStmt
parent <nil>
elmnts: [
	token switch switch

	node *ast.BadStmt
	parent *ast.SwitchStmt
	elmnts: []

	node *ast.Ident
	parent *ast.SwitchStmt
	elmnts: [
		token tag IDENT
	]

	node *ast.BlockStmt
	parent *ast.SwitchStmt
	elmnts: [
		token { {
	
		token } }
	]
]
`
	n := &ast.SwitchStmt{
		Switch: token.Pos(1),
		Init:   &ast.BadStmt{},
		Tag:    getIdent("tag"),
		Body:   &ast.BlockStmt{},
	}
	checkSyntaxTree(t, e, n)
}

func TestTypeSwitchStmtNode(t *testing.T) {
	e := `node *ast.TypeSwitchStmt
parent <nil>
elmnts: [
	token switch switch

	node *ast.BadStmt
	parent *ast.TypeSwitchStmt
	elmnts: []

	node *ast.BadStmt
	parent *ast.TypeSwitchStmt
	elmnts: []

	node *ast.BlockStmt
	parent *ast.TypeSwitchStmt
	elmnts: [
		token { {
	
		token } }
	]
]
`
	n := &ast.TypeSwitchStmt{
		Switch: token.Pos(1),
		Init:   &ast.BadStmt{},
		Assign: &ast.BadStmt{},
		Body:   &ast.BlockStmt{},
	}
	checkSyntaxTree(t, e, n)
}

func TestCommClauseNode(t *testing.T) {
	e := `node *ast.CommClause
parent <nil>
elmnts: [
	token case case

	node *ast.BadStmt
	parent *ast.CommClause
	elmnts: []

	token : :

	node *ast.BadStmt
	parent *ast.CommClause
	elmnts: []
]
`
	n := &ast.CommClause{
		Case:  token.Pos(1),
		Comm:  &ast.BadStmt{},
		Colon: token.Pos(2),
		Body:  []ast.Stmt{&ast.BadStmt{}},
	}

	checkSyntaxTree(t, e, n)
}

func TestSelectStmtNode(t *testing.T) {
	e := `node *ast.SelectStmt
parent <nil>
elmnts: [
	token select select

	node *ast.BlockStmt
	parent *ast.SelectStmt
	elmnts: [
		token { {
	
		token } }
	]
]
`
	n := &ast.SelectStmt{
		Select: token.Pos(1),
		Body:   &ast.BlockStmt{},
	}
	checkSyntaxTree(t, e, n)
}

func TestForStmtNode(t *testing.T) {
	e := `node *ast.ForStmt
parent <nil>
elmnts: [
	token for for

	node *ast.BadStmt
	parent *ast.ForStmt
	elmnts: []

	node *ast.Ident
	parent *ast.ForStmt
	elmnts: [
		token cond IDENT
	]

	node *ast.BadStmt
	parent *ast.ForStmt
	elmnts: []

	node *ast.BlockStmt
	parent *ast.ForStmt
	elmnts: [
		token { {
	
		token } }
	]
]
`
	n := &ast.ForStmt{
		For:  token.Pos(1),
		Init: &ast.BadStmt{},
		Cond: getIdent("cond"),
		Post: &ast.BadStmt{},
		Body: &ast.BlockStmt{},
	}
	checkSyntaxTree(t, e, n)
}

func TestRangeStmtNode(t *testing.T) {
	e := `node *ast.RangeStmt
parent <nil>
elmnts: [
	token for for

	node *ast.Ident
	parent *ast.RangeStmt
	elmnts: [
		token key IDENT
	]

	node *ast.Ident
	parent *ast.RangeStmt
	elmnts: [
		token value IDENT
	]

	token := :=

	node *ast.Ident
	parent *ast.RangeStmt
	elmnts: [
		token x IDENT
	]

	node *ast.BlockStmt
	parent *ast.RangeStmt
	elmnts: [
		token { {
	
		token } }
	]
]
`
	n := &ast.RangeStmt{
		For:    token.Pos(1),
		Key:    getIdent("key"),
		Value:  getIdent("value"),
		TokPos: token.Pos(2),
		Tok:    token.DEFINE,
		X:      getIdent("x"),
		Body:   &ast.BlockStmt{},
	}
	checkSyntaxTree(t, e, n)
}

func TestImportSpecNode(t *testing.T) {
	e := `node *ast.ImportSpec
parent <nil>
elmnts: [
	node *ast.CommentGroup
	parent *ast.ImportSpec
	elmnts: [
		node *ast.Comment
		parent *ast.CommentGroup
		elmnts: []
	]

	node *ast.Ident
	parent *ast.ImportSpec
	elmnts: [
		token id IDENT
	]

	node *ast.BasicLit
	parent *ast.ImportSpec
	elmnts: [
		token path STRING
	]

	node *ast.CommentGroup
	parent *ast.ImportSpec
	elmnts: [
		node *ast.Comment
		parent *ast.CommentGroup
		elmnts: []
	]
]
`
	n := &ast.ImportSpec{
		Doc:     getCommentGroup("c"),
		Name:    getIdent("id"),
		Path:    getBasicLit(token.STRING, "path"),
		Comment: getCommentGroup("c2"),
		EndPos:  token.Pos(1),
	}
	checkSyntaxTree(t, e, n)
}
