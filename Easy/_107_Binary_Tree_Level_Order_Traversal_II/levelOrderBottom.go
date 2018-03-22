package _107_Binary_Tree_Level_Order_Traversal_II

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

var ret [][]int

func levelTravel(nodeArr []*TreeNode){
	var nextNodeArr []*TreeNode
	var valArr []int

	len1 := len(nodeArr)
	if 0 == len1 {
		return
	}

	for i := 0; i < len1; i++ {
		v := nodeArr[i]
		if nil != v.Left {
			nextNodeArr = append(nextNodeArr, v.Left)
		}
		if nil != v.Right {
			nextNodeArr = append(nextNodeArr, v.Right)
		}

		valArr = append(valArr, v.Val)
	}
	fmt.Printf("valArr:%+v\n", valArr)
	levelTravel(nextNodeArr)

	ret = append(ret, valArr)
}

func levelOrderBottom(root *TreeNode) [][]int {
	if nil == root {
		return [][]int{}
	}
	ret = [][]int{}

	nextNodeArr := []*TreeNode{root}
	levelTravel(nextNodeArr)

	return ret
}
