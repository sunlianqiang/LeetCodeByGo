package _198_House_Robber

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

func isSymmetric(root *TreeNode) bool {
	if nil == root {
		return true
	}

	nextNodeArr := []*TreeNode{root}
	var notAllNil bool
	notAllNil = true
	for ;len(nextNodeArr) > 0 && notAllNil; {
		notAllNil = false
		len1 := len(nextNodeArr)

		var tmpArr []*TreeNode
		i, j := 0, len1-1
		for ; i < j;i, j = i+1, j-1 {

			if nextNodeArr[i] != nextNodeArr[j] {
				return false
			}

			if nil != nextNodeArr[i].Left {
				tmpArr = append(tmpArr, nextNodeArr[i].Left)
				notAllNil = true
			} else {
				tmpArr = append(tmpArr, nil)
			}

			if nil != nextNodeArr[i].Right {
				tmpArr = append(tmpArr, nextNodeArr[i].Right)
				notAllNil = true
			} else {
				tmpArr = append(tmpArr, nil)
			}

		}

		for ; i <  len1; i++ {
			if nil != nextNodeArr[i].Left {
				tmpArr = append(tmpArr, nextNodeArr[i].Left)
				notAllNil = true
			} else {
				tmpArr = append(tmpArr, nil)
			}

			if nil != nextNodeArr[i].Right {
				tmpArr = append(tmpArr, nextNodeArr[i].Right)
				notAllNil = true
			} else {
				tmpArr = append(tmpArr, nil)
			}
		}

		nextNodeArr = tmpArr

	}

	return false
}

