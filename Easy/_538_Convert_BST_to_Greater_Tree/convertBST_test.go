package _538_Convert_BST_to_Greater_Tree

import (
	"testing"
	"fmt"
)

func InitTree(index int, input []int) (t *TreeNode) {
	if index > len(input) {
		return nil
	}

	t = new(TreeNode)
	t.Val = input[index - 1]

	t.Left = InitTree(index*2, input)
	t.Right = InitTree(index*2 + 1, input)

	return t
}

func MiddleOrder(t *TreeNode) {
	if nil == t {
		return
	}

	MiddleOrder(t.Left)
	fmt.Printf("%+v, ", t.Val)
	MiddleOrder(t.Right)
}

func TestConvertBST(t *testing.T) {

	input1 := []int{5, 2, 13}
	tree1 := InitTree(1, input1)

	input2 := []int{2,1,3}
	tree2 := InitTree(1, input2)

	var tests = []struct{
		root *TreeNode
	}{
		{tree1,},
		{tree2,},
	}

	for _, v := range tests {
		MiddleOrder(v.root)
		fmt.Println()
		ConvertBST(v.root)

		MiddleOrder(v.root)
		fmt.Println()
	}


}