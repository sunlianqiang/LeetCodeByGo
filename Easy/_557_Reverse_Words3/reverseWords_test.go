package _557_Reverse_Words3

import "testing"

func TestReverseWords(t *testing.T) {
	ret := ReverseWords("Let's take LeetCode contest")
	want := "s'teL ekat edoCteeL tsetnoc"

	if ret == want {
		t.Logf("pass")
	} else {
		t.Errorf("fail, want %+v, get %+v", want, ret)
	}
}