package _485_Max_Consecutive_Ones

func FindMaxConsecutiveOnes(nums []int) int {
	var maxConsecutiveOnes int
	length := len(nums)
	for i := 0; i < length; {
		if 0 == nums[i] {
			i++
			continue
		} else {
			consecutiveOnes := 1

			for i++; i< length && 1 == nums[i]; i++ {
					consecutiveOnes++
			}
			if consecutiveOnes > maxConsecutiveOnes {
				maxConsecutiveOnes = consecutiveOnes
			}
		}
	}

	return maxConsecutiveOnes
}
