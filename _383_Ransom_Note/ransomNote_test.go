package _383_Ransom_Note

import "testing"

func TestCanConstruct(t *testing.T) {
	var tests = []struct{
		a string
		b string
		output bool
	} {
		{"aa", "ab", false},
	}

	for _, test := range tests {
		ret := CanConstruct(test.a, test.b)

		if ret == test.output {
			t.Logf("pass")
		} else {
			t.Errorf("fail, want %+v, get %+v", test.output, ret)
		}
	}
}
