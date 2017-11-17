package _506_Relative_Ranks

import "testing"

func TestFindRelativeRanks(t *testing.T) {
	var tests = []struct{
		input []int
		output []string
	}{
		{[]int{10,3,8,9,4}, []string{"Gold Medal","5","Bronze Medal","Silver Medal","4"}},
		{[]int{}, []string{}},
		{[]int{3}, []string{"Gold Medal"}},
		{[]int{3, 2, 1}, []string{"Gold Medal","Silver Medal","Bronze Medal"}},
		{[]int{5,4,3,2,1}, []string{"Gold Medal","Silver Medal","Bronze Medal","4","5"}},
	}

	for _, test := range tests {
		ret := FindRelativeRanks(test.input)

		t.Logf("want:%+v\n, get:%+v\n", test.output, ret)
	}
}
