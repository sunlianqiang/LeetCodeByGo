package _594_Longest_Harmonious_Subsequence

func findLHS(nums []int) int {
	var numMap map[int]int
	numMap = make(map[int]int)

	for _, v := range nums {
		numMap[v]++
	}

	var LHS int
	for k, v := range numMap {
		tmp1, ok1 := numMap[k-1]
		if ok1 && tmp1 + v > LHS {
			LHS = tmp1 + v
		}

		tmp2, ok2 := numMap[k+1]
		if ok2 && tmp2 + v > LHS {
			LHS = tmp2 + v
		}
	}

	return LHS
}
