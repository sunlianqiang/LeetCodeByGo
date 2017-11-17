package _26_Remove_Duplicates_from_Sorted_Array

import "fmt"

func RemoveDuplicates(nums []int) int {
	len1 := len(nums)
	if 0 == len1 {
		return 0
	}

	i, j := 1, 0
	for ; i < len1; i++{
		if nums[i] > nums[j]{
			j++
			nums[j] = nums[i]
		}

	}
	j++
	fmt.Printf("nums:%+v\n", nums)
	return j
}
