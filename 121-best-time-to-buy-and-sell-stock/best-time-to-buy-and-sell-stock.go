func maxProfit(prices []int) int {
	lowest := 10_001
	ans := 0
	for _, price := range prices {
		if price < lowest {
			lowest = price
			continue
		}
		profit := price - lowest
		if profit > ans {
			ans = profit
		}
	}

	return ans
}