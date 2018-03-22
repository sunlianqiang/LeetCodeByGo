package _485_Max_Consecutive_Ones

import "testing"

func TestFindMaxConsecutiveOnes(t *testing.T) {
	input := []int{1,1,0,1,1,1}
	want := 3

	ret := FindMaxConsecutiveOnes(input)

	if want == ret {
		t.Logf("pass")
	} else {
		t.Errorf("fail, want %+v, get %+v", want, ret)
	}
}
