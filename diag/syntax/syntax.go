package syntax

import (
	"go/ast"
	"go/token"
	"reflect"
)

// ElementType denotes element type
type ElementType int

// Syntax element types
const (
	ElementTypeToken ElementType = iota
	ElementTypeNode
)

// Element represents node or token
type Element interface {
	GetParent() Node
	GetElementType() ElementType
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
	GetText() string
	GetKind() token.Token
}

// FromAstNode returns node from the given ast.Node
func FromAstNode(node ast.Node) Node {
	return fromAstNodeAndParent(nil, node)
}

// IsNode returns true if element is syntax node
func IsNode(elmt Element) bool {
	if elmt == nil {
		return false
	}
	return elmt.GetElementType() == ElementTypeNode
}

// IsToken returns true if element is syntax token
func IsToken(elmt Element) bool {
	if elmt == nil {
		return false
	}
	return elmt.GetElementType() == ElementTypeToken
}

// GetElementTypeString returns element type string
func GetElementTypeString(elmt Element) string {
	if IsNode(elmt) {
		return "node"
	}

	if IsToken(elmt) {
		return "token"
	}

	return "none"
}

type nodeImpl struct {
	Parent   Node
	AstNode  ast.Node
	Elements []Element
}

func (n *nodeImpl) GetElementType() ElementType {
	return ElementTypeNode
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

func (t *tokenImpl) GetElementType() ElementType {
	return ElementTypeToken
}

func (t *tokenImpl) GetParent() Node {
	return t.Parent
}

func (t *tokenImpl) GetText() string {
	return t.Text
}

func (t *tokenImpl) GetKind() token.Token {
	return t.Kind
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

func isNilNode(node ast.Node) bool {
	return node == nil || reflect.ValueOf(node).IsNil()
}

func appendToken(elmts []Element, parent Node, pos token.Pos, text string, kind token.Token) []Element {
	elmt := tokenNode(parent, pos, text, kind)
	return append(elmts, elmt)
}

func appendLParenToken(elmts []Element, parent Node, pos token.Pos) []Element {
	return appendToken(elmts, parent, pos, "(", token.LPAREN)
}

func appendRParenToken(elmts []Element, parent Node, pos token.Pos) []Element {
	return appendToken(elmts, parent, pos, ")", token.RPAREN)
}

func appendLBraceToken(elmts []Element, parent Node, pos token.Pos) []Element {
	return appendToken(elmts, parent, pos, "{", token.LBRACE)
}

func appendRBraceToken(elmts []Element, parent Node, pos token.Pos) []Element {
	return appendToken(elmts, parent, pos, "}", token.RBRACE)
}

func appendLBrackToken(elmts []Element, parent Node, pos token.Pos) []Element {
	return appendToken(elmts, parent, pos, "[", token.LBRACK)
}

func appendRBrackToken(elmts []Element, parent Node, pos token.Pos) []Element {
	return appendToken(elmts, parent, pos, "]", token.RBRACK)
}

func appendElement(elmts []Element, parent Node, node ast.Node) []Element {
	if isNilNode(node) {
		return elmts
	}
	elmt := fromAstNodeAndParent(parent, node)
	return append(elmts, elmt)
}

func appendIdents(elmts []Element, parent Node, idents []*ast.Ident) []Element {
	if idents == nil {
		return elmts
	}
	for _, ident := range idents {
		elmts = appendElement(elmts, parent, ident)
	}
	return elmts
}

func appendComments(elmts []Element, parent Node, comments []*ast.Comment) []Element {
	if comments == nil {
		return elmts
	}
	for _, c := range comments {
		elmts = appendElement(elmts, parent, c)
	}
	return elmts
}

func appendFields(elmts []Element, parent Node, fields []*ast.Field) []Element {
	if fields == nil {
		return elmts
	}
	for _, f := range fields {
		elmts = appendElement(elmts, parent, f)
	}
	return elmts
}

func appendStmts(elmts []Element, parent Node, stmts []ast.Stmt) []Element {
	if stmts == nil {
		return elmts
	}
	for _, s := range stmts {
		elmts = appendElement(elmts, parent, s)
	}
	return elmts
}

func appendExprs(elmts []Element, parent Node, exprs []ast.Expr) []Element {
	if exprs == nil {
		return elmts
	}
	for _, e := range exprs {
		elmts = appendElement(elmts, parent, e)
	}
	return elmts
}

func loadElements(parent Node, node ast.Node) []Element {
	elmts := []Element{}

	switch n := node.(type) {
	case *ast.Comment:
		return nil

	case *ast.CommentGroup:
		elmts = appendComments(elmts, parent, n.List)
		return elmts

	case *ast.Field:
		elmts = appendElement(elmts, parent, n.Doc)
		elmts = appendIdents(elmts, parent, n.Names)
		elmts = appendElement(elmts, parent, n.Type)
		elmts = appendElement(elmts, parent, n.Tag)
		elmts = appendElement(elmts, parent, n.Comment)
		return elmts

	case *ast.FieldList:
		elmts = appendLParenToken(elmts, parent, n.Opening)
		elmts = appendFields(elmts, parent, n.List)
		elmts = appendRParenToken(elmts, parent, n.Closing)
		return elmts

	case *ast.BadExpr:
		return nil

	case *ast.Ident:
		elmts = appendToken(elmts, parent, n.NamePos, n.Name, token.IDENT)
		return elmts

	case *ast.Ellipsis:
		elmts = appendToken(elmts, parent, n.Ellipsis, "...", token.ELLIPSIS)
		elmts = appendElement(elmts, parent, n.Elt)
		return elmts

	case *ast.BasicLit:
		elmts = appendToken(elmts, parent, n.ValuePos, n.Value, n.Kind)
		return elmts

	case *ast.FuncLit:
		elmts = appendElement(elmts, parent, n.Type)
		elmts = appendElement(elmts, parent, n.Body)
		return elmts

	case *ast.CompositeLit:
		elmts = appendElement(elmts, parent, n.Type)
		elmts = appendLBraceToken(elmts, parent, n.Lbrace)
		elmts = appendExprs(elmts, parent, n.Elts)
		elmts = appendRBraceToken(elmts, parent, n.Rbrace)
		return elmts

	case *ast.ParenExpr:
		elmts = appendLParenToken(elmts, parent, n.Lparen)
		elmts = appendElement(elmts, parent, n.X)
		elmts = appendRParenToken(elmts, parent, n.Rparen)
		return elmts

	case *ast.SelectorExpr:
		elmts = appendElement(elmts, parent, n.X)
		elmts = appendElement(elmts, parent, n.Sel)
		return elmts

	case *ast.IndexExpr:
		elmts = appendElement(elmts, parent, n.X)
		elmts = appendLBrackToken(elmts, parent, n.Lbrack)
		elmts = appendElement(elmts, parent, n.Index)
		elmts = appendRBrackToken(elmts, parent, n.Rbrack)
		return elmts

	case *ast.SliceExpr:
		elmts = appendElement(elmts, parent, n.X)
		elmts = appendLBrackToken(elmts, parent, n.Lbrack)
		elmts = appendElement(elmts, parent, n.Low)
		elmts = appendElement(elmts, parent, n.High)
		elmts = appendElement(elmts, parent, n.Max)
		elmts = appendRBrackToken(elmts, parent, n.Rbrack)
		return elmts

	case *ast.TypeAssertExpr:
		elmts = appendElement(elmts, parent, n.X)
		elmts = appendLParenToken(elmts, parent, n.Lparen)
		elmts = appendElement(elmts, parent, n.Type)
		elmts = appendRParenToken(elmts, parent, n.Rparen)
		return elmts

	case *ast.CallExpr:
		elmts = appendElement(elmts, parent, n.Fun)
		elmts = appendLParenToken(elmts, parent, n.Lparen)
		elmts = appendExprs(elmts, parent, n.Args)
		elmts = appendToken(elmts, parent, n.Ellipsis, "...", token.ELLIPSIS)
		elmts = appendRParenToken(elmts, parent, n.Rparen)
		return elmts

	case *ast.StarExpr:
		elmts = appendToken(elmts, parent, n.Star, "*", token.MUL)
		elmts = appendElement(elmts, parent, n.X)
		return elmts

	case *ast.UnaryExpr:
		elmts = appendToken(elmts, parent, n.OpPos, n.Op.String(), n.Op)
		elmts = appendElement(elmts, parent, n.X)
		return elmts

	case *ast.BinaryExpr:
		elmts = appendElement(elmts, parent, n.X)
		elmts = appendToken(elmts, parent, n.OpPos, n.Op.String(), n.Op)
		elmts = appendElement(elmts, parent, n.Y)
		return elmts

	case *ast.KeyValueExpr:
		elmts = appendElement(elmts, parent, n.Key)
		elmts = appendToken(elmts, parent, n.Colon, ":", token.COLON)
		elmts = appendElement(elmts, parent, n.Value)
		return elmts

	case *ast.ArrayType:
		elmts = appendLBrackToken(elmts, parent, n.Lbrack)
		elmts = appendElement(elmts, parent, n.Len)
		elmts = appendElement(elmts, parent, n.Elt)
		return elmts

	case *ast.StructType:
		elmts = appendToken(elmts, parent, n.Struct, "struct", token.STRUCT)
		elmts = appendElement(elmts, parent, n.Fields)
		return elmts

	case *ast.FuncType:
		elmts = appendToken(elmts, parent, n.Func, "func", token.FUNC)
		elmts = appendElement(elmts, parent, n.Params)
		elmts = appendElement(elmts, parent, n.Results)
		return elmts

	case *ast.InterfaceType:
		elmts = appendToken(elmts, parent, n.Interface, "interface", token.INTERFACE)
		elmts = appendElement(elmts, parent, n.Methods)
		return elmts

	case *ast.MapType:
		elmts = appendToken(elmts, parent, n.Map, "map", token.MAP)
		elmts = appendElement(elmts, parent, n.Key)
		elmts = appendElement(elmts, parent, n.Value)
		return elmts

	case *ast.ChanType:
		if n.Begin != n.Arrow {
			elmts = appendToken(elmts, parent, n.Begin, "chan", token.CHAN)
			elmts = appendToken(elmts, parent, n.Arrow, "<-", token.ARROW)
		} else {
			elmts = appendToken(elmts, parent, n.Arrow, "<-", token.ARROW)
		}
		elmts = appendElement(elmts, parent, n.Value)
		return elmts

	case *ast.BadStmt:
		return nil

	case *ast.DeclStmt:
		elmts = appendElement(elmts, parent, n.Decl)
		return elmts

	case *ast.EmptyStmt:
		if !n.Implicit {
			elmts = appendToken(elmts, parent, n.Semicolon, ";", token.SEMICOLON)
			return elmts
		}
		return nil

	case *ast.LabeledStmt:
		elmts = appendElement(elmts, parent, n.Label)
		elmts = appendToken(elmts, parent, n.Colon, ":", token.COLON)
		elmts = appendElement(elmts, parent, n.Stmt)
		return elmts

	case *ast.ExprStmt:
		elmts = appendElement(elmts, parent, n.X)
		return elmts

	case *ast.SendStmt:
		elmts = appendElement(elmts, parent, n.Chan)
		elmts = appendToken(elmts, parent, n.Arrow, "<-", token.ARROW)
		elmts = appendElement(elmts, parent, n.Value)
		return elmts

	case *ast.IncDecStmt:
		elmts = appendElement(elmts, parent, n.X)
		elmts = appendToken(elmts, parent, n.TokPos, n.Tok.String(), n.Tok)
		return elmts

	case *ast.AssignStmt:
		elmts = appendExprs(elmts, parent, n.Lhs)
		elmts = appendToken(elmts, parent, n.TokPos, n.Tok.String(), n.Tok)
		elmts = appendExprs(elmts, parent, n.Rhs)
		return elmts

	case *ast.GoStmt:
		elmts = appendToken(elmts, parent, n.Go, "go", token.GO)
		elmts = appendElement(elmts, parent, n.Call)
		return elmts

	case *ast.DeferStmt:
		elmts = appendToken(elmts, parent, n.Defer, "defer", token.DEFER)
		elmts = appendElement(elmts, parent, n.Call)
		return elmts

	case *ast.ReturnStmt:
		elmts = appendToken(elmts, parent, n.Return, "return", token.RETURN)
		elmts = appendExprs(elmts, parent, n.Results)
		return elmts

	case *ast.BranchStmt:
		elmts = appendToken(elmts, parent, n.TokPos, n.Tok.String(), n.Tok)
		elmts = appendElement(elmts, parent, n.Label)
		return elmts

	case *ast.BlockStmt:
		elmts = appendLBraceToken(elmts, parent, n.Lbrace)
		elmts = appendStmts(elmts, parent, n.List)
		elmts = appendRBraceToken(elmts, parent, n.Rbrace)
		return elmts

	case *ast.IfStmt:
		elmts = appendToken(elmts, parent, n.If, token.IF.String(), token.IF)
		elmts = appendElement(elmts, parent, n.Init)
		elmts = appendElement(elmts, parent, n.Cond)
		elmts = appendElement(elmts, parent, n.Body)
		elmts = appendElement(elmts, parent, n.Else)
		return elmts

	case *ast.CaseClause:
		elmts = appendToken(elmts, parent, n.Case, token.CASE.String(), token.CASE)
		elmts = appendExprs(elmts, parent, n.List)
		elmts = appendToken(elmts, parent, n.Colon, token.COLON.String(), token.COLON)
		elmts = appendStmts(elmts, parent, n.Body)
		return elmts

	case *ast.SwitchStmt:
		elmts = appendToken(elmts, parent, n.Switch, token.SWITCH.String(), token.SWITCH)
		elmts = appendElement(elmts, parent, n.Init)
		elmts = appendElement(elmts, parent, n.Tag)
		elmts = appendElement(elmts, parent, n.Body)
		return elmts

	case *ast.TypeSwitchStmt:
		elmts = appendToken(elmts, parent, n.Switch, token.SWITCH.String(), token.SWITCH)
		elmts = appendElement(elmts, parent, n.Init)
		elmts = appendElement(elmts, parent, n.Assign)
		elmts = appendElement(elmts, parent, n.Body)
		return elmts

	case *ast.CommClause:
		elmts = appendToken(elmts, parent, n.Case, token.CASE.String(), token.CASE)
		elmts = appendElement(elmts, parent, n.Comm)
		elmts = appendToken(elmts, parent, n.Colon, token.COLON.String(), token.COLON)
		elmts = appendStmts(elmts, parent, n.Body)
		return elmts

	case *ast.SelectStmt:
		elmts = appendToken(elmts, parent, n.Select, token.SELECT.String(), token.SELECT)
		elmts = appendElement(elmts, parent, n.Body)
		return elmts

	case *ast.ForStmt:
		elmts = appendToken(elmts, parent, n.For, token.FOR.String(), token.FOR)
		elmts = appendElement(elmts, parent, n.Init)
		elmts = appendElement(elmts, parent, n.Cond)
		elmts = appendElement(elmts, parent, n.Post)
		elmts = appendElement(elmts, parent, n.Body)
		return elmts

	case *ast.RangeStmt:
		elmts = appendToken(elmts, parent, n.For, token.FOR.String(), token.FOR)
		elmts = appendElement(elmts, parent, n.Key)
		elmts = appendElement(elmts, parent, n.Value)
		elmts = appendToken(elmts, parent, n.TokPos, n.Tok.String(), n.Tok)
		elmts = appendElement(elmts, parent, n.X)
		elmts = appendElement(elmts, parent, n.Body)
		return elmts
	}
	return nil
}
