package asttest

import (
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/a6cexz/goanalyzer/diag/text"
	"github.com/a6cexz/goanalyzer/diag/text/helpers"
	"golang.org/x/tools/go/ast/astutil"
)

// ParseTestFile parses test file
func ParseTestFile(src string) (*token.FileSet, *ast.File, text.TextSpan, error) {
	src, m := helpers.RemoveTextMarkers(src)
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "source.go", src, parser.ParseComments)
	if err != nil {
		return fset, file, text.NewTextSpan(0, 0), err
	}

	start := 0
	if pos, ok := m["#start#"]; ok {
		start = pos
	}

	end := start
	if pos, ok := m["#end#"]; ok {
		end = pos
	}

	s := text.NewTextSpanFromBounds(start, end)
	return fset, file, s, nil
}

// GetTestAstNode gets test node from source
func GetTestAstNode(src string) (ast.Node, error) {
	_, file, span, err := ParseTestFile(src)
	if err != nil {
		return nil, err
	}

	start := token.Pos(span.Start())
	end := token.Pos(span.End())
	path, _ := astutil.PathEnclosingInterval(file, start, end)
	node := path[0]

	return node, nil
}
