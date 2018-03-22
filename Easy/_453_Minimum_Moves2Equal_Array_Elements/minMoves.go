package _453_Minimum_Moves2Equal_Array_Elements

func MinMoves(nums []int) int {
	var min int
	min = nums[0]
	length := len(nums)
	for i := 1; i < length; i++ {
		if min > nums[i] {
			min = nums[i]
		}
	}
	if 1 == length {
		return 0
	}
	var sum int
	for i := 0; i < length; i++ {
		sum += nums[i]
	}

	return sum - min * length
}
