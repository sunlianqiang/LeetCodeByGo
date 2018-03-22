package _530_Minimum_Absolute_Difference_in_BST

import "fmt"

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

var nodeSlice []int

func PreOrder(t *TreeNode)  {
	if nil == t {
		return
	}

	PreOrder(t.Left)
	nodeSlice = append(nodeSlice, t.Val)
	PreOrder(t.Right)
}
func GetMinimumDifference(root *TreeNode) int {
	nodeSlice = []int{}
	PreOrder(root)

	fmt.Printf("slice:%+v", nodeSlice)

	len1 := len(nodeSlice)

	minDiff := nodeSlice[1] - nodeSlice[0]
	for i := 2; i < len1; i++ {
		diff := nodeSlice[i] - nodeSlice[i-1]
		if diff < minDiff {
			minDiff = diff
		}
	}
	return minDiff
}