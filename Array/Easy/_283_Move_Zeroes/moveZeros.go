package _283_Move_Zeroes


func MoveZeroes(nums []int)  {
	length := len(nums)
	var zeroSlice []int
	for i := 0; i < length; i++ {
		if 0 == nums[i] {
			zeroSlice = append(zeroSlice, 0)
			nums = append(nums[0:i], nums[i+1:]...)
			i--
			length--
		}
	}

	nums = append(nums, zeroSlice...)
}