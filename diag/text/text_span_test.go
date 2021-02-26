package text_test

import (
	"testing"

	"github.com/a6cexz/goanalyzer/diag/text"
	"github.com/stretchr/testify/assert"
)

func Test_TextSpan_NewTextSpanFromBounds(t *testing.T) {
	s := text.NewTextSpanFromBounds(0, 5)
	assert.Equal(t, 0, s.Start())
	assert.Equal(t, 5, s.End())
	assert.Equal(t, 5, s.Length())
}

func Test_TextSpan_IsValid(t *testing.T) {
	assert.True(t, text.NewTextSpan(0, 0).IsValid())
	assert.True(t, text.NewTextSpan(0, 1).IsValid())
	assert.False(t, text.NewTextSpan(-1, 1).IsValid())
	assert.False(t, text.NewTextSpan(0, -1).IsValid())
	assert.False(t, text.NewTextSpan(10, -1).IsValid())
	assert.True(t, text.NewTextSpan(10, 0).IsValid())
}

func Test_TextSpan_Start_End_Length_IsEmpty(t *testing.T) {
	s := text.NewTextSpan(0, 10)
	assert.Equal(t, 0, s.Start())
	assert.Equal(t, 10, s.Length())
	assert.Equal(t, 10, s.End())
	assert.False(t, s.IsEmpty())
}

func Test_TextSpan_String(t *testing.T) {
	s := text.NewTextSpan(0, 10)
	assert.Equal(t, "[0..10)", s.String())
}

func Test_TextSpan_CompareTo(t *testing.T) {
	s1 := text.NewTextSpan(0, 10)
	s2 := text.NewTextSpan(0, 10)
	assert.True(t, s1.CompareTo(s2) == 0)
	assert.True(t, s2.CompareTo(s1) == 0)

	s1 = text.NewTextSpan(0, 10)
	s2 = text.NewTextSpan(1, 10)
	assert.True(t, s1.CompareTo(s2) < 0)
	assert.True(t, s2.CompareTo(s1) > 0)

	s1 = text.NewTextSpan(0, 10)
	s2 = text.NewTextSpan(0, 11)
	assert.True(t, s1.CompareTo(s2) < 0)
	assert.True(t, s2.CompareTo(s1) > 0)
}

func Test_TextSpan_ContainsPos(t *testing.T) {
	s := text.NewTextSpan(0, 4)
	assert.True(t, s.ContainsPos(-1))
	assert.True(t, s.ContainsPos(0))
	assert.True(t, s.ContainsPos(1))
	assert.True(t, s.ContainsPos(2))
	assert.True(t, s.ContainsPos(3))
	assert.False(t, s.ContainsPos(4))
	assert.False(t, s.ContainsPos(5))
	assert.False(t, s.ContainsPos(6))
}

func Test_TextSpan_ContainsSpan(t *testing.T) {
	s1 := text.NewTextSpan(0, 4)
	s2 := text.NewTextSpan(0, 4)
	s3 := text.NewTextSpan(0, 3)
	s4 := text.NewTextSpan(1, 3)
	s5 := text.NewTextSpan(0, 5)
	assert.True(t, s1.ContainsSpan(s2))
	assert.True(t, s1.ContainsSpan(s3))
	assert.True(t, s1.ContainsSpan(s4))
	assert.False(t, s1.ContainsSpan(s5))
	assert.True(t, s5.ContainsSpan(s1))
}

func Test_TextSpan_Equals(t *testing.T) {
	s1 := text.NewTextSpan(0, 4)
	s2 := text.NewTextSpan(0, 4)
	s3 := text.NewTextSpan(0, 3)
	assert.True(t, s1.Equals(s2))
	assert.False(t, s1.Equals(s3))
}
