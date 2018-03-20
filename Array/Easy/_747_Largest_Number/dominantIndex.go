package _747_Largest_Number

func dominantIndex(nums []int) int {
	length := len(nums)
	if 1 == length {
		return 0
	}
	var ret int
	var max int
	max = nums[0]
	for i := 1; i < length; i++ {
		v := nums[i]
		if v > max {
			max = v
			ret = i
		}
	}

	for _, v := range nums {
		if v == max {
			continue
		}

		if max < v*2 {
			return -1
		}
	}

	return ret
}
