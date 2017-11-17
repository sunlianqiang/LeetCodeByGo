package _438_Find_All_Anagrams_in_a_String

import "testing"

func equalSlice(s, p []int) bool {
	len1 := len(s)
	len2 := len(p)
	if len1 != len2 {
		return false
	}

	for i := 0; i < len1; i++ {
		if s[i] != p[i] {
			return false
		}
	}

	return true
}

func TestFindAnagrams(t *testing.T) {
	var tests = []struct{
		s string
		p string
		ret []int
	} {
		{"cbaebabacd", "abc", []int{0,6}},
		{"baa", "aa", []int{1}},
	}

	for _, v := range tests {
		ret := FindAnagrams(v.s, v.p)
		if equalSlice(ret, v.ret) {
			t.Logf("pass")
		} else {
			t.Errorf("fail, want %+v, get %+v", v.ret, ret)
		}
	}
}
