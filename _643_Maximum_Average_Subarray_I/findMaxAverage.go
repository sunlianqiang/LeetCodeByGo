package _643_Maximum_Average_Subarray_I

func findMaxAverage(nums []int, k int) float64 {
	len1 := len(nums)
	var sum, maxSum int
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	maxSum = sum

	for j := k; j < len1; j++ {
		sum =  nums[j] - nums[j-k]
		if maxSum < sum {
			maxSum = sum
		}
	}

	return float64(maxSum) / float64(k)
}
