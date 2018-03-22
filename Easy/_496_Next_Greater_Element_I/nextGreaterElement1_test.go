package _496_Next_Greater_Element_I

import "testing"

func isSliceEqual(ret, want []int) bool {
	len1 := len(ret)
	len2 := len(want)

	if len1 != len2 {
		return false
	}

	for i := 0; i < len1; i++ {
		if ret[i] != want[i] {
			return false
		}
	}

	return true
}
func TestNextGreaterElement(t *testing.T) {
	nums1, nums2 := []int{4,1,2}, []int{1,3,4,2}
	want := []int{-1,3,-1}

	ret := NextGreaterElement(nums1, nums2)

	if isSliceEqual(ret, want) {
		t.Logf("pass")
	} else {
		t.Errorf("fail, want %+v, get %+v", want, ret)
	}
}

func TestNextGreaterElement2(t *testing.T) {
	nums1, nums2 := []int{2,4}, []int{1,2,3,4}
	want := []int{3,-1}

	ret := NextGreaterElement(nums1, nums2)

	if isSliceEqual(ret, want) {
		t.Logf("pass")
	} else {
		t.Errorf("fail, want %+v, get %+v", want, ret)
	}
}