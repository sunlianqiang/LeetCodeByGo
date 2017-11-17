package _504_Base7

import "testing"

func TestConvertToBase7(t *testing.T) {
	var tests = []struct{
		input int
		output string
	} {
		{-7, "-10"},
		{101, "203"},
		{0, "0"},
	}

	for _, v := range tests {
		ret := ConvertToBase7(v.input)

		if ret == v.output {
			t.Logf("pass")
		} else {
			t.Errorf("fail, want %+v, get %+v\n", v.output, ret)
		}
	}
}
