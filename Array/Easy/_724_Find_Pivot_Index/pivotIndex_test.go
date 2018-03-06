package _724_Find_Pivot_Index

import (
	"testing"
)

func Test_PivotIndex(t *testing.T) {
	nums := []int{1, 7, 3, 6, 5, 6}
	ret := pivotIndex(nums)

	want := 3
	if want == ret {
		t.Logf("pass")
	} else {
		t.Errorf("fail, want %+v, get %+v", want, ret)
	}
}
