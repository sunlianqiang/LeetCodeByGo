package _387_First_Unique_Character_String

import "testing"

func TestFirstUniqChar(t *testing.T) {
	var tests = []struct{
		intput string
		output int
	}{
		{"loveleetcode", 2},
		{"", -1},
		{"cc", -1},
	}

	for _, test := range tests {
		ret := FirstUniqChar(test.intput)

		if ret == test.output {
			t.Logf("pass")
		} else {
			t.Errorf("fail, want %+v, get %+v", test.output, ret)
		}
	}
}
