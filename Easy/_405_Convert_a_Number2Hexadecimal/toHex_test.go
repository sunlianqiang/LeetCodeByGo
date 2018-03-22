package _405_Convert_a_Number2Hexadecimal

import "testing"

func TestToHex(t *testing.T) {
	var tests = []struct{
		num int
		out string
	}{
		{0, "0"},
		{26, "1a"},
		{-1, "ffffffff"},
	}

	for _, v := range tests {
		ret := ToHex(v.num)

		if ret == v.out {
			t.Logf("pass")
		} else {
			t.Errorf("fail, want %+v,  get %+v\n", v.out, ret)
		}
	}
}
