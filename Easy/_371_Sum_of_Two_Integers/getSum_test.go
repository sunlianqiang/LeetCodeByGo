package _371_Sum_of_Two_Integers

import "testing"

func TestGetSum(t *testing.T) {
	var tests = []struct {
		a int
		b int
		output int
	}{
		{1, 2, 3},
		{1, -1, 0},
	}

	for _, test := range tests {
		ret := GetSum(test.a, test.b)
		if ret == test.output {
			t.Logf("pass")
		} else {
			t.Errorf("fail, want %+v, get %+v", test.output, ret)
		}
	}
}
