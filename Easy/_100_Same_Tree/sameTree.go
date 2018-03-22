package _100_Same_Tree

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

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if nil == p && nil == q {
		return true
	} else if nil == p {
		return false
	} else if nil == q {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	if false == isSameTree(p.Left, q.Left) {
		return false
	}
	if false == isSameTree(p.Right, q.Right) {
		return false
	}

	return true
}
