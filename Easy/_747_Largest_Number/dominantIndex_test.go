package _747_Largest_Number

import (
	"testing"
)

type Input struct {
	nums     []int
	expected int
}

func TestDominatIndex(t *testing.T) {
	var inputs = []Input{
		Input{
			nums:     []int{0, 0, 2, 3},
			expected: -1,
		},
	}

	for _, input := range inputs {
		ret := dominantIndex(input.nums)

		if ret == input.expected {
			t.Logf("Pass")
		} else {
			t.Errorf("Fail, expect %v, get %v\n", input.expected, ret)
		}
	}
}
