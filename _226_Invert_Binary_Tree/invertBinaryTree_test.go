package _226_Invert_Binary_Tree

import (
	"testing"
	"fmt"
)

func InitNode(val int, left *TreeNode, right *TreeNode) (ret *TreeNode){
	ret = new(TreeNode)
	ret.Val = val
	ret.Left = left
	ret.Right = right

	return ret
}


func treeEqual(t1 *TreeNode, t2 *TreeNode) bool{
	if t1.Val != t2.Val {
		return false
	}

	if nil == t1.Left && nil != t2.Left {
		return false
	}else if nil != t1.Left && nil == t2.Left {
		return false
	} else if nil != t1.Left && nil != t2.Left {
		left := treeEqual(t1.Left, t2.Left)
		if !left {
			return false
		}
	}

	if nil == t1.Right && nil != t2.Right {
		return false
	}else if nil != t1.Right && nil == t2.Right {
		return false
	}else if nil != t1.Right && nil != t2.Right {
		right := treeEqual(t1.Right, t2.Right)
		if !right {
			return false
		}
	}

	return true
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

func TestInvertTree(t *testing.T) {
	l3_1 := InitNode(1, nil, nil)
	l3_2 := InitNode(3, nil, nil)
	l3_3 := InitNode(6, nil, nil)
	l3_4 := InitNode(9, nil, nil)

	l2_1 := InitNode(2, l3_1, l3_2)
	l2_2 := InitNode(7, l3_3, l3_4)

	input := InitNode(4, l2_1, l2_2)



	w2_1 := InitNode(7, l3_4, l3_3)
	w2_2 := InitNode(2, l3_2, l3_1)

	want := InitNode(4, w2_1, w2_2)

	ret := InvertTree(input)

	equal := treeEqual(want, ret)

	TreePrint(want)
	fmt.Println()
	TreePrint(ret)
	fmt.Println()

	if equal {
		t.Logf("pass")
	} else {
		t.Errorf("fail, want %+v, get %+v", want, ret)
	}
}
