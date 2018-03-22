package _53_Maximum_Subarray

func maxSubArray(nums []int) int {
	len1 := len(nums)
	if 0 == len1 {
		return 0
	}

	preSum := nums[0]
	maxSum := preSum
	for i := 1; i < len1; i++ {
		if preSum < 0 {
			preSum = nums[i]
		} else {
			preSum = nums[i] + preSum
		}

		if maxSum < preSum {
			maxSum = preSum
		}
	}

	return maxSum
}
