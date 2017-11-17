package _543_Diameter_Binary_Tree

import "math"

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
var maxDiameter int
func MaxDeepth(t *TreeNode) (deepth int)  {
	if nil == t {
		return 0
	}

	left := MaxDeepth(t.Left)
	right := MaxDeepth(t.Right)

	maxDiameter = int(math.Max(float64(maxDiameter), float64(left + right)))

	return  int(math.Max(float64(left), float64(right))) + 1
}


func DiameterOfBinaryTree(root *TreeNode) int {
	maxDiameter = 0
	MaxDeepth(root)

	return maxDiameter
}
