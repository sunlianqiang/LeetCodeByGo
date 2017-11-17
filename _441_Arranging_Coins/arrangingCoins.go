package _441_Arranging_Coins

func ArrangeCoins(n int) int {
	if 0 == n {
		return 0
	}

	var ret int
	for i := 0; i <= n; i++ {
		if i * (i+1)/2 <= n && n < (i+1)*(i+2)/2 {
			ret = i
			break
		}
	}

	return ret
}
