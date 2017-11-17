package _645_Set_Mismatch

func findErrorNums(nums []int) []int {
	var sum int
	var repeatNum int
	var missNum int
	var numMap map[int]int
	numMap = make(map[int]int)

	for _, v := range nums {
		numMap[v]++
		if 2 == numMap[v] {
			repeatNum = v
		}

		sum += v
	}

	len1 := len(nums)
	realSum := (len1 + 1) * len1 /2
	missNum = repeatNum + (realSum - sum)

	return []int{repeatNum, missNum}
}
