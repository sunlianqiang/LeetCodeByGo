package _637_Average_Levels_Binary_Tree

import "testing"


func InitNode(val int, left *TreeNode, right *TreeNode) (ret *TreeNode){
	ret = new(TreeNode)
	ret.Val = val
	ret.Left = left
	ret.Right = right

	return ret
}

func SliceEqual(s1, s2 []float64) bool {
	len1 := len(s1)
	len2 := len(s2)

	if len1 != len2 {
		return false
	}
	for i := 0; i < len1; i++ {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}
func TestAverageOfLevels(t *testing.T) {
	l3_1 := InitNode(15, nil, nil)
	l3_2 := InitNode(7, nil, nil)

	l2_1 := InitNode(9, nil, nil)
	l2_2 := InitNode(20, l3_1, l3_2)

	l1 := InitNode(3, l2_1, l2_2)

	ret := AverageOfLevels(l1)

	want := []float64{3, 14.5, 11}

	ok := SliceEqual(ret, want)

	if ok {
		t.Logf("pass")
	} else {
		t.Errorf("fail, want %+v, get %+v", want, ret)
	}
}