package _724_Find_Pivot_Index

import (
	"testing"
)

type Input struct {
	nums []int
	want int
}

var inputs = []Input{
	Input{
		nums: []int{1, 7, 3, 6, 5, 6},
		want: 3,
	},
	Input{
		nums: []int{1, 2, 3},
		want: -1,
	},
	Input{
		nums: []int{},
		want: -1,
	},
	Input{
		nums: []int{3},
		want: 0,
	}, Input{
		nums: []int{-1, -1, -1, -1, -1, 0},
		want: 2,
	},
}

func Test_PivotIndex1(t *testing.T) {
	for _, input := range inputs {
		nums := input.nums
		t.Logf("input:%v\n", nums)
		ret := pivotIndex(nums)

		want := input.want
		if want == ret {
			t.Logf("pass")
		} else {
			t.Errorf("fail, want %+v, get %+v", want, ret)
			return
		}
	}

}

func Test_PivotIndex2(t *testing.T) {
	nums := inputs[4].nums
	t.Logf("input:%v\n", nums)
	ret := pivotIndex(nums)

	want := 2
	if want == ret {
		t.Logf("pass")
	} else {
		t.Errorf("fail, want %+v, get %+v", want, ret)
	}
}
