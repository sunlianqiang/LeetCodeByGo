package _653_Two_Sum_IV_Input_BST

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

//middle search
func MiddleSearch(root *TreeNode, valSlice *[]int) {
	if nil == root {
		return
	}
	MiddleSearch(root.Left, valSlice)
	(*valSlice) = append(*valSlice, root.Val)
	MiddleSearch(root.Right, valSlice)
}
func FindTarget(root *TreeNode, k int) bool {
	if nil == root {
		return false
	}

	var valSlice []int
	MiddleSearch(root, &valSlice)

	fmt.Printf("val:%+v", valSlice)
	for i, j := 0, len(valSlice) - 1; i < j; {
		if k == valSlice[i] + valSlice[j] {
			return true
		} else if k < valSlice[i] + valSlice[j] {
			j--
		} else if k > valSlice[i] + valSlice[j] {
			i++
		}

	}

	return false
}
