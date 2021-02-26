package text

import (
	"fmt"
)

// TextSpan represents span in text
type TextSpan struct {
	start  int
	length int
}

// NewTextSpan creates new text span
func NewTextSpan(start, length int) TextSpan {
	return TextSpan{start: start, length: length}
}

// NewTextSpanFromBounds creates new text span
func NewTextSpanFromBounds(start, end int) TextSpan {
	if start < 0 {
		panic("start is negative!")
	}

	if end < start {
		panic("end is less than start!")
	}

	return NewTextSpan(start, end-start)
}

// Start returns text span start
func (s TextSpan) Start() int {
	return s.start
}

// End returns text span end
func (s TextSpan) End() int {
	return s.start + s.length
}

// Length returns text span length
func (s TextSpan) Length() int {
	return s.length
}

// IsValid checks if text span is valid
func (s TextSpan) IsValid() bool {
	if s.start < 0 {
		return false
	}

	if s.start+s.length < s.start {
		return false
	}

	return true
}

// IsEmpty returns true if span length is zero
func (s TextSpan) IsEmpty() bool {
	return s.length == 0
}

// String returns string representation of the text span
func (s TextSpan) String() string {
	return fmt.Sprintf("[%d..%d)", s.Start(), s.End())
}

// CompareTo compares spans
func (s TextSpan) CompareTo(span TextSpan) int {
	r := s.start - span.start
	if r != 0 {
		return r
	}
	return s.length - span.length
}

// ContainsPos checks if position is insider the span
func (s TextSpan) ContainsPos(pos int) bool {
	return pos-s.start < s.length
}

// ContainsSpan checks if span contains another span
func (s TextSpan) ContainsSpan(span TextSpan) bool {
	return span.start >= s.start && span.End() <= s.End()
}

// Equals checks if span equals to another span
func (s TextSpan) Equals(span TextSpan) bool {
	return s.start == span.start && s.length == span.length
}
