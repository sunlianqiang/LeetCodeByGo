package _563_Binary_Tree_Tilt

import "testing"

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

func TestFindTilt(t *testing.T) {
	input1 := []int{1,2,3}

	var tests = []struct{
		tree *TreeNode
		output int
	}{
		{InitTree(1, input1), 1},
	}

	for _, test := range tests {
		ret := FindTilt(test.tree)

		if ret == test.output {
			t.Logf("pass")
		} else {
			t.Errorf("fail, want %+v, get %+v", test.output, ret)
		}
	}
}