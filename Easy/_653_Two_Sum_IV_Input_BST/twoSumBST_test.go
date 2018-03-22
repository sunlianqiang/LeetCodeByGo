package _653_Two_Sum_IV_Input_BST

import "testing"

func InitNode(val int, left *TreeNode, right *TreeNode) (ret *TreeNode){
	ret = new(TreeNode)
	ret.Val = val
	ret.Left = left
	ret.Right = right

	return ret
}

func TestFindTarget(t *testing.T) {
	l3_1 := InitNode(2, nil, nil)
	l3_2 := InitNode(4, nil, nil)
	l3_3 := InitNode(7, nil, nil)

	l2_1 := InitNode(3, l3_1, l3_2)
	l2_2 := InitNode(6, nil, l3_3)
	root := InitNode(5, l2_1, l2_2)

	want := true

	ret := FindTarget(root, 9)

	if ret == want {
		t.Logf("pass")
	} else {
		t.Errorf("fail, want %+v, get %+v", want, ret)
	}

	r2_1 := InitNode(1, nil, nil)
	r2_2 := InitNode(3, nil, nil)

	input2 := InitNode(2, r2_1, r2_2)

	want2 := true
	ret2 := FindTarget(input2, 3)

	if ret2 == want2 {
		t.Logf("pass")
	} else {
		t.Errorf("fail, want %+v, get %+v", want2, ret2)
	}

}
