package syntax_test

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestBadExprNode(t *testing.T) {
	e := `node *ast.BadExpr
parent <nil>
elmnts: []
`
	expr := getBadExpr()
	checkSyntaxTree2(t, e, expr)
}

func TestIdentNode(t *testing.T) {
	e := `node *ast.Ident
parent <nil>
elmnts: [
	token test IDENT
]
`
	n := getIdent("test")
	checkSyntaxTree2(t, e, n)
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
	checkSyntaxTree2(t, e, n)
}

func TestBasicLitNode(t *testing.T) {
	e := `node *ast.BasicLit
parent <nil>
elmnts: [
	token test STRING
]
`
	n := getBasicLit(token.STRING, "test")
	checkSyntaxTree2(t, e, n)
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

	checkSyntaxTree2(t, e, n)
}
