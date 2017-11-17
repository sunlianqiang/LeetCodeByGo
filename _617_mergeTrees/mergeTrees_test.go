package _617_mergeTrees

import (
	"testing"
	"fmt"
)


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

func Test_mergeTrees(t *testing.T) {
	t1l2 := InitNode(5, nil, nil)
	t1l1 := InitNode(3, t1l2, nil)
	t1r1 := InitNode(2, nil, nil)
	t1 := InitNode(1, t1l1, t1r1)

	t2l2r1 := InitNode(4, nil, nil)
	t2l2r2 := InitNode(7, nil, nil)
	t2l1r1 := InitNode(1, nil, t2l2r1)
	t2l1r2 := InitNode(3, nil, t2l2r2)
	t2 := InitNode(2, t2l1r1, t2l1r2)

	t3l2r1 := InitNode(5, nil, nil)
	t3l2r2 := InitNode(4, nil, nil)
	t3l2r3 := InitNode(7, nil, nil)

	t3l1r1 := InitNode(4, t3l2r1, t3l2r2)
	t3l1r2 := InitNode(5, nil, t3l2r3)
	t3 := InitNode(3, t3l1r1, t3l1r2)

	fmt.Printf("merge t1:")
	TreePrint(t1)
	fmt.Println()

	fmt.Printf("merge t2:")
	TreePrint(t2)
	fmt.Println()

	ret := MergeTrees(t1, t2)

	ok := treeEqual(t3, ret)

	fmt.Printf("merge t3:")
	TreePrint(t1)
	fmt.Println()

	fmt.Printf("t1:%+v, t2:%+v, t3:%+v\n", t1, t2, t3)
	if !ok {
		t.Errorf("fail, ret.Val want %+v, get %+v\n", t3, ret)
	} else {
		t.Logf("pass")
	}



}
