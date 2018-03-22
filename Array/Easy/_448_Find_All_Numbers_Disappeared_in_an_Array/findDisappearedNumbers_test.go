package _448_Find_All_Numbers_Disappeared_in_an_Array

import "testing"

func sliceEqual(want, ret []int) bool {
	len1 := len(want)
	len2 := len(ret)

	if len1 != len2 {
		return false
	}

	for i := 0; i < len1; i++ {
		if want[i] != ret[i] {
			return false
		}
	}

	return true
}

func TestFindDisappearedNumbers(t *testing.T) {
	input := []int{4,3,2,7,8,2,3,1}
	want := []int{5,6}

	ret := FindDisappearedNumbers(input)

	ok := sliceEqual(want, ret)
	if ok {
		t.Logf("pass")
	} else {
		t.Errorf("fail, want %+v, get %+v", want, ret)
	}
}
