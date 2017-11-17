package _404_Sum_Left_Leaves

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

func PreOrderTravel(t *TreeNode, left bool) {
	if nil == t {
		return
	}
	if left && nil == t.Left && nil == t.Right {
		sum += t.Val
	}
	PreOrderTravel(t.Left, true)
	PreOrderTravel(t.Right, false)

}
var sum int
func sumOfLeftLeaves(root *TreeNode) int {
	sum = 0
	PreOrderTravel(root, false)

	return sum
}
