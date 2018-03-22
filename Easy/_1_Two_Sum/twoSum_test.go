package _1_Two_Sum

import "testing"

func TestTwoSum(t *testing.T) {
	var tests = []struct{
		input []int
		target int
		output []int
	}{
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}

	for _, v := range tests {
		ret := TwoSum(v.input, v.target)

		if ret[0] == v.output[0] && ret[1] == v.output[1] {
			t.Logf("pass")
		} else if ret[0] == v.output[1] && ret[1] == v.output[0] {
			t.Logf("pass")
		} else {
			t.Errorf("fail, want %+v, get %+v", v.output, ret)
		}
	}
}
