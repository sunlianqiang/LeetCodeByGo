package _606_Construct_String_from_Binary_Tree

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

func PreOrder(t *TreeNode, ret *string) {
	if nil == t {
		return
	}
	(*ret) += fmt.Sprintf("(%+v", t.Val)

	if nil == t.Left && nil != t.Right {
		(*ret) += "()"
	}

	PreOrder(t.Left, ret)
	PreOrder(t.Right, ret)
	(*ret) += ")"
}

func Tree2str(t *TreeNode) string {
	if nil == t {
		return ""
	}

	var ret string
	ret += fmt.Sprintf("%+v", t.Val)
	if nil == t.Left && nil != t.Right {
		ret += "()"
	}
	PreOrder(t.Left, &ret)
	PreOrder(t.Right, &ret)


	return ret
}