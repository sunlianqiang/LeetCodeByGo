package _205_Isomorphic_Strings

import "testing"

func TestIsIsomorphic(t *testing.T) {
	var tests = []struct{
		s string
		t string
		output bool
	} {
		{"egg", "add", true},
		{"foo", "bar", false},
		{"paper", "title", true},
		{"ab", "aa", false},
		{"aa", "ab", false},
		{"aba", "baa", false},
	}

	for _, v := range tests {
		ret := IsIsomorphic(v.s, v.t)
		if ret == v.output {
			t.Logf("pass")
		} else {
			t.Errorf("fail, want %+v, get %+v", v.output, ret)
		}
	}
}
