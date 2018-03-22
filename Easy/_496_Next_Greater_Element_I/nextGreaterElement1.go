package _496_Next_Greater_Element_I


func NextGreaterElement(findNums []int, nums []int) []int {
	var ret []int
	var numsMap map[int]int

	numsLen := len(nums)

	numsMap = make(map[int]int, numsLen)

	for i := 0; i < numsLen; i++ {
		numsMap[nums[i]] = i
	}


	findNumsLen := len(findNums)
	for i := 0; i < findNumsLen; i++ {
		num := findNums[i]
		pos, _ := numsMap[num]

		var ok bool
		for j := pos; j < numsLen; j++ {
			if nums[j] > num {
				ret = append(ret, nums[j])
				ok = true
				break
			}
		}
		if !ok {
			ret = append(ret, -1)
		}
	}

	return ret
}