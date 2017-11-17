package _70_Climbing_Stairs

func climbStairs(n int) int {
	if 1 == n {
		return 1
	} else if 2 == n {
		return 2
	}


	ret1 := 1
	ret2 := 2
	var ret int
	for i := 2; i < n; i++ {
		ret = ret1 + ret2
		ret1 = ret2
		ret2 = ret
	}

	return ret
}
