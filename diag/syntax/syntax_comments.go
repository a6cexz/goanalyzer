package syntax

import "go/ast"

// Comment node
type Comment struct {
	*nodeImpl
}

func newComment(parent Node, node *ast.Comment) *Comment {
	if node == nil {
		return nil
	}
	r := &Comment{}
	r.nodeImpl = getNodeImpl(parent, node)
	return r
}

func newComments(parent Node, nodes []*ast.Comment) []*Comment {
	if nodes == nil {
		return nil
	}
	comments := []*Comment{}
	for _, c := range nodes {
		comment := newComment(parent, c)
		comments = append(comments, comment)
	}
	return comments
}

// CommentGroup node
type CommentGroup struct {
	*nodeImpl
	List []*Comment
}

func newCommentGroup(parent Node, node *ast.CommentGroup) *CommentGroup {
	if node == nil {
		return nil
	}
	r := &CommentGroup{}
	r.nodeImpl = getNodeImpl(parent, node)
	r.List = newComments(r, node.List)
	r.Elements = getElements(r)
	return r
}
