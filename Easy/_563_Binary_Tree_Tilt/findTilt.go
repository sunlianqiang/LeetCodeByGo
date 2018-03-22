package _563_Binary_Tree_Tilt

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

var tilt int
func subTreeSum(t *TreeNode) (sum int){
	if nil == t {
		return 0
	}

	left := subTreeSum(t.Left)
	right := subTreeSum(t.Right)

	tilt += int(math.Abs((float64(left) - float64(right))))

	return t.Val + left + right
}
func FindTilt(root *TreeNode) int {
	tilt = 0
	subTreeSum(root)
	return tilt
}
