package _283_Move_Zeroes

import "testing"

func SliceEqual(a, b []int) bool{
	lena := len(a)
	lenb := len(b)

	if lena != lenb {
		return false
	}

	for i := 0; i < lena; i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestMoveZeroes(t *testing.T) {
	var tests = []struct{
		input []int
		output []int
	}{
		{[]int{0, 1, 0, 3, 12},
			[]int{1, 3, 12, 0, 0}},
	}

	for _, test := range tests {
		MoveZeroes(test.input)
		ret := test.input
		ok := SliceEqual(ret, test.output)
		if ok {
			t.Logf("pass")
		} else {
			t.Errorf("fail, want %+v, get %+v", test.output, ret)
		}
	}
}
