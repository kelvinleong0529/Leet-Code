func maxProfit(prices []int) int {
	lowest := 10_001
	highest := -1
	ans := 0
	for _, price := range prices {
		if price < lowest {
			lowest = price
			highest = -1
			continue
		}
		if price > highest {
			highest = price
			profit := highest - lowest
			if profit > ans {
				ans = profit
			}
		}
	}

	return ans
}