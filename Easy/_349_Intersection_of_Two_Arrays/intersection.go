package _349_Intersection_of_Two_Arrays

func intersection(nums1 []int, nums2 []int) []int {
	var ret []int

	var numMap map[int]bool
	numMap = make(map[int]bool)

	for _, v := range nums1 {
		numMap[v] = false
	}

	for _, v := range nums2 {
		_, ok := numMap[v]
		if ok {
			numMap[v] = true
		}
	}

	for k, v := range numMap {
		if v  {
			ret = append(ret, k)
		}
	}
	return ret
}