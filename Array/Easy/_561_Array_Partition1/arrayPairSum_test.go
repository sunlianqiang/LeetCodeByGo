package _561_Array_Partition1

import "testing"

func Test_ArrayPairSum(t *testing.T) {

	input := []int{1,4,3,2}

	sum := ArrayPairSum(input)

	if 4 == sum {
		t.Logf("pass")
	} else {
		t.Errorf("fail, want 4, get %+v\n", sum)
	}
}