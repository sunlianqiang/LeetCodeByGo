package _101_Symmetric_Tree

import "golang.org/x/text/unicode/bidi"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func ReverseTree(t *TreeNode) {
	if nil == t {
		return
	}
	t.Left, t.Right = t.Right, t.Left
	ReverseTree(t.Left)
	ReverseTree(t.Right)
}

func EqualTree(t1, t2 *TreeNode) bool  {
	if nil == t1 && nil == t2 {
		return true
	} else if nil != t1 && nil == t2 {
		return false
	} else if nil == t1 && nil != t2 {
		return false
	}

	if t1.Val != t2.Val {
		return false
	}

	ok1 := EqualTree(t1.Left, t2.Left)
	if !ok1 {
		return false
	}

	ok2 := EqualTree(t1.Right, t2.Right)
	if !ok2 {
		return false
	}

	return true
}

func isSymmetric(root *TreeNode) bool {
	if nil == root {
		return true
	}

	left := root.Left
	ReverseTree(root.Right)

	equal := EqualTree(left, root.Right)

	if equal {
		return true
	} else {
		return false
	}
}
