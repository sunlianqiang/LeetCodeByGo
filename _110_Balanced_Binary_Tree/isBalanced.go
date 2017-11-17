package _110_Balanced_Binary_Tree

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

func Height(t *TreeNode) (h int)  {
	if nil == t {
		return 0
	}
	left := Height(t.Left)
	right := Height(t.Right)

	if left == 0 && right == 0 {
		return 1
	}

	if left > right {
		return left + 1
	} else {
		return right + 1
	}

}

func isBalanced(root *TreeNode) bool {
	if nil == root {
		return true
	}

	diff := Height(root.Left) - Height(root.Right)
	if diff < -1 || diff > 1 {
		return false
	}

	if isBalanced(root.Left)  && isBalanced(root.Right) {
		return true
	}

	return false
}
