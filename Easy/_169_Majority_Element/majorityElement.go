package _169_Majority_Element

func majorityElement(nums []int) int {
	var ret int
	var majorityMap map[int]int
	majorityMap = make(map[int]int)

	len1 := len(nums)
	for i := 0; i < len1; i++{
		num := nums[i]
		majorityMap[num]++

		if majorityMap[num] > len1/2 {
			ret = num
			break
		}
	}

	return ret
}
