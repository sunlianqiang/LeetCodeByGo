package _258_Add_Digits

import "testing"

func TestAddDigits(t *testing.T) {
	var tests = []struct{
		num int
		output int
	}{
		{38, 2},
		{10, 1},
	}

	for _, test := range tests {
		ret := AddDigits(test.num)

		if ret == test.output {
			t.Logf("pass")
		} else {
			t.Errorf("fail, want %+v, get %+v", test.output, ret)
		}
	}
}
