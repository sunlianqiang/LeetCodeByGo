package _344_Reverse_String

import "testing"

func TestReverseString(t *testing.T) {
	input := "hello"
	want := "olleh"

	ret := ReverseString(input)

	if ret == want {
		t.Logf("pass")
	} else {
		t.Errorf("fail, want %+v, get %+v", want, ret)
	}
}
