package _26_Remove_Duplicates_from_Sorted_Array

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	var tests = []struct{
		input []int
		ret int
	} {
		{[]int{1, 1}, 1},
		{[]int{1,1,1}, 1},
		{[]int{1,2,2,3}, 3},
		{[]int{-999,-999,-998,-998,-997,-997}, 3},
		}

	for _, v := range tests {
		ret := RemoveDuplicates(v.input)

		if ret == v.ret {
			t.Logf("pass")
		} else {
			t.Errorf("fail, want %+v, get %+v", v.ret, ret)
		}
	}

}