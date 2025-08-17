func maxProfit(prices []int) int {
	lowest := 10_001
	highest := -1
	ans := 0
	for _, price := range prices {
		// found a better entry opportunity
		// immediately sell the previous stock
		if price < lowest {
			ans += calProfit(lowest, highest)
			lowest = price
			highest = -1
			continue
		}
		// if current price is lower than the highest
		// sell the current stock and take profit
		// buy at the current price
		if price < highest {
			ans += calProfit(lowest, highest)
			lowest = price
			highest = -1
			continue
		}
		if price > highest {
			highest = price
		}
	}

	return ans + calProfit(lowest, highest)
}

func calProfit(lowest, highest int) int {
	if highest == -1 {
		return 0
	}
	return highest - lowest
}