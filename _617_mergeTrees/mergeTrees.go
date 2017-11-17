package _617_mergeTrees

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
	 Val int
     Left *TreeNode
     Right *TreeNode
 }

func InitNode(val int, left *TreeNode, right *TreeNode) (ret *TreeNode){
	ret = new(TreeNode)
	ret.Val = val
	ret.Left = left
	ret.Right = right

	return ret
}

func TreePrint(t1 *TreeNode) {

	if nil == t1 {
		fmt.Printf("null, ")
		return
	} else {
		fmt.Printf("%+v, ", t1.Val)
	}

	TreePrint(t1.Left)
	TreePrint(t1.Right)
}


func MergeTrees(t1 *TreeNode, t2 *TreeNode) (t3 *TreeNode) {

	if nil == t1 && nil == t2 {
		return nil
	}

	if nil != t1 && nil == t2{
		return t1
	}else if nil == t1 && nil != t2 {
		return t2
	} else {
		t1.Val += t2.Val
		t1.Left = MergeTrees(t1.Left, t2.Left)
		t1.Right = MergeTrees(t1.Right, t2.Right)
	}

	return t1
}