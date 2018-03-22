package _572_Subtree_of_Another_Tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isTreeEqual(s *TreeNode, t *TreeNode) bool {
	if nil == s && nil == t {
		return true
	} else if nil != s && nil == t {
		return false
	} else if nil == s && nil != t {
		return false
	}

	if s.Val != t.Val {
		return false
	}
	if !isTreeEqual(s.Left, t.Left) {
		return false
	}
	if !isTreeEqual(s.Right, t.Right) {
		return false
	}

	return true
}

var subTree bool

func travelTree(s *TreeNode, t *TreeNode)  {
	//already find subtree
	if subTree {
		return
	}

	if nil == s {
		return
	}

	if isTreeEqual(s, t) {
		subTree = true
		return
	}

	travelTree(s.Left, t)
	travelTree(s.Right, t)
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	subTree = false

	travelTree(s, t)
	return subTree
}