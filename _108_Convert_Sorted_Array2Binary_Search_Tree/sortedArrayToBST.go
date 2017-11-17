package _108_Convert_Sorted_Array2Binary_Search_Tree

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

func InitBST(input []int) (t *TreeNode) {
	len1 := len(input)

	if 0 == len1 {
		return nil
	}

	middle := len1 / 2
	t = new(TreeNode)
	t.Val = input[middle]

	t.Left = InitBST(input[:middle])
	t.Right = InitBST(input[middle+1:])

	return t
}

func sortedArrayToBST(nums []int) *TreeNode {
	var root *TreeNode

	root = InitBST(nums)

	return root
}
