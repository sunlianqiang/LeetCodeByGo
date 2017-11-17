package _521_Longest_Uncommon_Subsequence_I

import "testing"

func TestFindLUSlength(t *testing.T) {
	input1 := "aba"
	input2 := "cdc"

	want := 3
	ret := FindLUSlength(input1, input2)

	if want == ret {
		t.Logf("pass")
	} else {
		t.Errorf("fail, want %+v, get %+v", want, ret)
	}
}