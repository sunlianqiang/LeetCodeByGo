package _350_Intersection_Two_Arrays_II

import "testing"

func TestIntersect(t *testing.T) {
	var tests = []struct{
		input1 []int
		input2 []int
	} {
		{[]int{1}, []int{1}, },
	}

	for _, v := range tests {
		Intersect(v.input1, v.input2)
	}
}
