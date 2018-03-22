package _606_Construct_String_from_Binary_Tree

import "testing"

func InitNode(val int, left *TreeNode, right *TreeNode) (ret *TreeNode){
	ret = new(TreeNode)
	ret.Val = val
	ret.Left = left
	ret.Right = right

	return ret
}

func TestTree2str(t *testing.T) {
	l3_1 := InitNode(4, nil, nil)

	l2_1 := InitNode(2, l3_1, nil)
	l2_2 := InitNode(3, nil, nil)

	input := InitNode(1, l2_1, l2_2)

	ret1 := Tree2str(input)
	want1 := "1(2(4))(3)"
	if ret1 == want1 {
		t.Logf("pass")
	} else {
		t.Errorf("fail, want %+v, get %+v", want1, ret1)
	}

	r3_1 := InitNode(4, nil, nil)

	r2_1 := InitNode(2, nil, r3_1)
	r2_2 := InitNode(3, nil, nil)

	input2 := InitNode(1, r2_1, r2_2)

	ret2 := Tree2str(input2)

	want2 := "1(2()(4))(3)"
	if ret2 == want2 {
		t.Logf("pass")
	} else {
		t.Errorf("fail, want %+v, get %+v", want2, ret2)
	}


	w6 := InitNode(6, nil, nil)
	w5 := InitNode(5, nil, w6)
	w4 := InitNode(4, nil, w5)
	w3 := InitNode(3, nil, w4)
	w2 := InitNode(2, nil, w3)
	input3 := InitNode(1, nil, w2)

	ret3 := Tree2str(input3)
	want3 := "1()(2()(3()(4()(5()(6)))))"
	if ret3 == want3 {
		t.Logf("pass")
	} else {
		t.Errorf("fail, want %+v, get %+v", want3, ret3)
	}


}
