package _226_Invert_Binary_Tree

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

func InvertTree(root *TreeNode) *TreeNode {
	if nil == root {
		return nil
	}

	if nil != root.Left {
		InvertTree(root.Left)
	}

	if nil != root.Right {
		InvertTree(root.Right)
	}

	root.Left, root.Right = root.Right, root.Left

	return root
}
