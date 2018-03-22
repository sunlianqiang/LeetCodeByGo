package _538_Convert_BST_to_Greater_Tree

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

var sum int
func ReverseOrderAdd(t *TreeNode, add int){
	if nil == t {
		return
	}

	ReverseOrderAdd(t.Right, add)
	t.Val += sum
	sum = t.Val
	ReverseOrderAdd(t.Left, sum)

}
func ConvertBST(root *TreeNode) *TreeNode {
	sum = 0
	ReverseOrderAdd(root, 0)

	return root
}
