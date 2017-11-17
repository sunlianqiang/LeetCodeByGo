package _350_Intersection_Two_Arrays_II

import "fmt"

func Intersect(nums1 []int, nums2 []int) []int {
	len1 := len(nums1)
	len2 := len(nums2)

	var numsShort []int
	var numsLong []int
	if len1 < len2 {
		numsShort = nums1
		numsLong = nums2
	} else {
		numsShort = nums2
		numsLong = nums1
	}

	fmt.Printf("short:%+v\n, long:%+v\n", numsShort, numsLong)
	var numMap map[int]int
	numMap = make(map[int]int)

	for _, v := range numsShort {
		numMap[v]++
	}

	fmt.Printf("map:%+v\n", numMap)
	var ret []int
	for _, v := range numsLong {
		tmp, ok := numMap[v]
		if ok && tmp > 0 {
			numMap[v]--
			ret = append(ret, v)
		}
	}

	return ret
}
