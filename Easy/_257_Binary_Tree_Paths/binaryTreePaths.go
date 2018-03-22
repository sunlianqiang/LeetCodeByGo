package _257_Binary_Tree_Paths

import "strconv"

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

var ret []string

func getPath(t *TreeNode, s string)  {
	val := t.Val
	s = s + "->" + strconv.Itoa(val)

	if nil == t.Left && nil == t.Right {
		ret = append(ret, s)
	}
	if nil != t.Left {
		getPath(t.Left, s)
	}
	if nil != t.Right {
		getPath(t.Right, s)
	}
}

func binaryTreePaths(root *TreeNode) []string {
	ret = []string{}
	if nil == root {
		return ret
	}

	rootval := root.Val

	if nil == root.Left && nil == root.Right {
		ret = append(ret, strconv.Itoa(rootval))
		return ret
	}
	if nil != root.Left {
		getPath(root.Left, strconv.Itoa(rootval))
	}
	if nil != root.Right {
		getPath(root.Right, strconv.Itoa(rootval))
	}


	return ret
}
