package _171_Excel_Sheet_Column_Number

import "testing"

func TestTitleToNumber(t *testing.T) {
	var tests = []struct{
		input string
		output int
	} {
		{"AB", 28},
		{"A", 1},
	}

	for _, test := range tests {
		ret := TitleToNumber(test.input)

		if ret == test.output {
			t.Logf("pass")
		} else {
			t.Errorf("fail, want %+v, get %+v", test.output, ret)
		}
	}
}
