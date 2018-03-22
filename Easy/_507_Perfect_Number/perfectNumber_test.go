package _507_Perfect_Number

import "testing"

func TestCheckPerfectNumber(t *testing.T) {
	var tests = []struct{
		num int
		ret bool
	} {
		{6, true},
		{28, true},
		{1, false},
	}

	for _, v := range tests {
		ret := CheckPerfectNumber(v.num)

		if ret == v.ret {
			t.Logf("pass")
		} else {
			t.Errorf("fail, want %+v, get %+v", v.ret, ret)
		}
	}
}
