package _501_Find_Mode_in_Binary_Search_Tree

import "uaiframework/utils/kafka"

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

var numMap map[int]int
func preOrderTravel(t *TreeNode)  {
	if nil == t {
		return
	}
	numMap[t.Val]++

	preOrderTravel(t.Left)
	preOrderTravel(t.Right)
}

func findMode(root *TreeNode) []int {
	numMap = make(map[int]int)
	preOrderTravel(root)


	type Mode struct {
		time int
		val int
	}
	var maxMode []Mode
	for k, v := range numMap {
		if v > maxMode[0].time {
			tmpMode := Mode{k, v}
			maxMode = []Mode{tmpMode}
		} else if v == maxMode[0].time {
			tmpMode := Mode{k, v}
			maxMode = append(maxMode, tmpMode)
		}
	}

	var ret []int
	for _, v := range maxMode {
		ret = append(ret, v.val)
	}
	return ret
}
