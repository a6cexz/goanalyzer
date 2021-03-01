package asttest_test

import (
	"testing"

	"github.com/a6cexz/goanalyzer/diag/syntax/asttest"
	"github.com/a6cexz/goanalyzer/diag/syntax/syntaxkind"
	"github.com/stretchr/testify/assert"
)

func TestParseTestFile(t *testing.T) {
	src := `package main
var a = 1#start#0
`
	node, err := asttest.GetTestAstNode(src)
	assert.NoError(t, err)
	assert.True(t, syntaxkind.IsExpr(node))
	assert.True(t, syntaxkind.AsExpr(node) != nil)
}
