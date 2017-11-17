package _35_Search_Insert_Position

import "testing"

func TestSearchInsert(t *testing.T) {
	var tests = []struct{
		input []int
		target int
		output int
	}{
		{[]int{1, 3, 5, 6}, 1, 0},
		{[]int{1, 3, 5, 6}, 3, 1},
		{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6}, 6, 3},
		{[]int{1, 3, 5, 6}, 0, 0},
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{1, 3, 5, 6}, 7, 4},
		{[]int{1, 3}, 2, 1},
	}

	for _, v := range tests {
		ret := SearchInsert(v.input, v.target)

		if ret == v.output {
			t.Logf("pass")
		} else {
			t.Errorf("fail, want %+v, get %+v\n", v.output, ret)
		}
	}
}
