package _724_Find_Pivot_Index

import "fmt"

func pivotIndex(nums []int) int {
	length := len(nums)
	if length < 1 {
		return -1
	} else if 1 == length {
		return 0
	}

	i := 0
	var totalLeft, totalRight int
	for _, num := range nums {
		totalRight += num
	}
	totalRight -= nums[0]

	for ; i < length; i++ {
		fmt.Printf("i:%v, totalLeft:%v, totalRight:%v\n", i, totalLeft, totalRight)
		if totalLeft == totalRight {
			return i
		}
		totalLeft += nums[i]
		if i < length-1 {
			totalRight -= nums[i+1]
		}

	}

	return -1
}
