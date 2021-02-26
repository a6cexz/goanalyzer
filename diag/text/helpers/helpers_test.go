package helpers_test

import (
	"testing"

	"github.com/a6cexz/goanalyzer/diag/text/helpers"
	"github.com/stretchr/testify/assert"
)

func TestGetTextMatches(t *testing.T) {
	src := `package main
var a = 1#tag_1#0
var b = 2#tag_2a#0
var c = 123#tag_3b#4
`
	e := `package main
var a = 10
var b = 20
var c = 1234
`
	s, m := helpers.RemoveTextMarkers(src)
	assert.Equal(t, e, s)
	assert.Equal(t, 22, m["#tag_1#"])
	assert.Equal(t, 33, m["#tag_2a#"])
	assert.Equal(t, 46, m["#tag_3b#"])
}
