package _437_Path_Sum_III

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

var ret int

func downTravel(t *TreeNode, sumNeed, sumGet int)  {
	if nil == t {
		return
	}

	sumGet += t.Val
	if sumGet == sumNeed {
		ret++
	}

	downTravel(t.Left, sumNeed, sumGet)
	downTravel(t.Right, sumNeed, sumGet)
}

func preOrderTravel(t *TreeNode, sum int)  {
	if nil == t {
		return
	}

	downTravel(t, sum, 0)

	preOrderTravel(t.Left, sum)
	preOrderTravel(t.Right, sum)
}

func pathSum(root *TreeNode, sum int) int {
	ret = 0
	preOrderTravel(root, sum)

	return ret
}