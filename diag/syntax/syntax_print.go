package syntax

import (
	"bytes"
	"fmt"
	"io"
)

// Print prints elmt to std output
func Print(elmt Element) {
	var buffer bytes.Buffer
	PrintTo(&buffer, elmt)
	str := buffer.String()
	fmt.Print(str)
}

// PrintTo prints elmt to the given writer
func PrintTo(w io.Writer, elmt Element) {
	printSyntaxRec(w, elmt, "")
}

func printSyntaxRec(w io.Writer, elmt Element, indent string) {
	if IsNode(elmt) {
		printNodeRec(w, elmt.(Node), indent)
		return
	}

	if IsToken(elmt) {
		printTokenRec(w, elmt.(Token), indent)
		return
	}
}

func printNodeRec(w io.Writer, node Node, indent string) {
	fmt.Fprintf(w, "%vnode %T\n", indent, node.GetAstNode())

	parent := node.GetParent()
	if parent != nil {
		fmt.Fprintf(w, "%vparent %T\n", indent, parent.GetAstNode())
	} else {
		fmt.Fprintf(w, "%vparent <nil>\n", indent)
	}

	elmnts := node.GetElements()
	count := len(elmnts)
	if count > 0 {
		fmt.Fprintf(w, "%velmnts: [\n", indent)
		for i, elmnt := range elmnts {
			printSyntaxRec(w, elmnt, "\t"+indent)
			if i < count-1 {
				fmt.Fprintf(w, "%v\n", indent)
			}
		}
		fmt.Fprintf(w, "%v]\n", indent)
	} else {
		fmt.Fprintf(w, "%velmnts: []\n", indent)
	}
}

func printTokenRec(w io.Writer, token Token, indent string) {
	fmt.Fprint(w, indent)
	fmt.Fprintf(w, "token %v %v\n", token.GetText(), token.GetKind())
}
