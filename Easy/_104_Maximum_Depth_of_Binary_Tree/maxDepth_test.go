package _104_Maximum_Depth_of_Binary_Tree

import "testing"

func InitNode(val int, left *TreeNode, right *TreeNode) (ret *TreeNode){
	ret = new(TreeNode)
	ret.Val = val
	ret.Left = left
	ret.Right = right

	return ret
}

func TestMaxDepth(t *testing.T) {
	l3_1 := InitNode(3, nil, nil)
	l2_1 := InitNode(2, l3_1, nil)
	root := InitNode(1, l2_1, nil)

	want := 3
	ret := MaxDepth(root)

	if ret == want {
		t.Logf("pass")
	} else {
		t.Errorf("fail, want %+v, get %+v", want, ret)
	}
}
