package _268_Missing_Number

func missingNumber(nums []int) int {
	len1 := len(nums)
	var numArr []bool
	numArr = make([]bool, len1 + 1)

	for i := 0; i < len1; i++ {
		numArr[nums[i]] = true
	}

	var ret int
	for k, v := range numArr {
		if false == v {
			ret = k
		}
	}

	return ret
}
