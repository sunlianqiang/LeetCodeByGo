package _121_Best_Time2Buy_and_Sell_Stock

func MaxProfit(prices []int) int {
	var maxProfit int
	var currentMin int
	len1 := len(prices)

	if len1 == 0 {
		return 0
	}
	currentMin = prices[0]


	for i := 1; i < len1; i++ {
		if maxProfit < prices[i] - currentMin {
			maxProfit = prices[i] - currentMin
		}

		if currentMin > prices[i] {
			currentMin = prices[i]
		}
	}

	return maxProfit
}
