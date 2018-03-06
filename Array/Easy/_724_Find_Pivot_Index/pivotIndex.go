package _724_Find_Pivot_Index

func pivotIndex(nums []int) int {
	i, j := 1, len(nums)-1
	var totalLeft, totalRight = nums[0], 0
	for i < j-1 {
		if totalLeft > totalRight {
			totalRight += nums[i+1]
			i++
		} else {
			totalLeft += nums[j-1]
			j--
		}
	}

	if totalLeft == totalRight {
		return i + 1
	} else {
		return -1
	}
}
