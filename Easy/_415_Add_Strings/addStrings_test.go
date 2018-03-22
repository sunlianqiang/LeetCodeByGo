package _415_Add_Strings

import "testing"

func TestAddStrings(t *testing.T) {
	var tests = []struct{
		num1 string
		num2 string
		output string
	}{
		{"0", "0", "0"},
		{"1", "2", "3"},
		{"123", "12", "135"},
	}

	for _, v := range tests {
		ret := AddStrings(v.num1, v.num2)

		if ret == v.output {
			t.Logf("pass")
		} else {
			t.Errorf("fail, want %+v, get %+v", v.output, ret)
		}
	}
}
