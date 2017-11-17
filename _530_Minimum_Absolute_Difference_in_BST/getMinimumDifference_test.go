package _530_Minimum_Absolute_Difference_in_BST

import "testing"


func InitTree(index int, input []int) (t *TreeNode) {
	if index > len(input) {
		return nil
	}

	if input[index - 1] < 0 {
		return nil
	}
	t = new(TreeNode)
	t.Val = input[index - 1]

	t.Left = InitTree(index*2, input)
	t.Right = InitTree(index*2 + 1, input)

	return t
}

func TestGetMinimumDifference(t *testing.T) {
	input1 := []int{1, -1, 3, -1, -1, 2}
	input2 := []int{1, -1, 5, -1, -1, 3}
	var tests = []struct{
		input *TreeNode
		output int
	}{
		{InitTree(1, input1), 1},
		{InitTree(1, input2), 2},
	}
	for _, test := range tests {
		ret := GetMinimumDifference(test.input)
		if ret == test.output {
			t.Logf("pass")
		} else {
			t.Errorf("fail, want %+v, get %+v", test.output, ret)
		}
	}

}
