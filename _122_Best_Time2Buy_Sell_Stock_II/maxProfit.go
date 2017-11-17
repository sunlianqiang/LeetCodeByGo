package _122_Best_Time2Buy_Sell_Stock_II

func maxProfit(prices []int) int {
	len1 := len(prices)

	var sum int

	for i := 1; i < len1; i++{
		if prices[i] > prices[i-1] {
			sum += prices[i] - prices[i-1]
		}

	}

	return sum
}
