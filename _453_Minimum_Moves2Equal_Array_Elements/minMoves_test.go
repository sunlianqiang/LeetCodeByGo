package _453_Minimum_Moves2Equal_Array_Elements

import "testing"

func TestMinMoves(t *testing.T) {
	var tests = []struct{
		input []int
		output int
	}{
		{
			[]int{1, 2, 3}, 3,
		},
		{
			[]int{83,86,77,15,93,35,86,92,49,2}, 598,
		},
	}

	for _, test := range tests {
		ret := MinMoves(test.input)

		if ret == test.output {
			t.Logf("pass")
		} else {
			t.Errorf("fail, want %+v, get %d", test.output, ret)
		}
	}
}
