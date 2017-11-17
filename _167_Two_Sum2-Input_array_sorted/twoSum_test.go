package _167_Two_Sum2_Input_array_sorted

import "testing"

func TestTwoSum(t *testing.T) {
	var tests = []struct{
		nums []int
		target int
		output []int
	}{
		{[]int{2,3,4}, 6, []int{1, 3}},
	}

	for _, test := range tests {
		ret := TwoSum(test.nums, test.target)

		if ret[0] == test.output[0] && ret[1] == test.output[1] {
			t.Logf("pass")
		} else {
			t.Errorf("fail, want %+v, get %+v", test.output, ret)
		}
	}
}