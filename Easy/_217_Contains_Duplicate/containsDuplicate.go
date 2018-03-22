package _217_Contains_Duplicate

func containsDuplicate(nums []int) bool {
	if len(nums) < 2 {
		return false
	}

	var numMap map[int]int
	numMap = make(map[int]int)

	for _, v := range nums {
		numMap[v]++

		if numMap[v] > 1 {
			return false
		}
	}

	return true
}