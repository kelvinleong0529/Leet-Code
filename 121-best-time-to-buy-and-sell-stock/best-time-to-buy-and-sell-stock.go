func maxProfit(prices []int) int {

	ans := 0
	buy := prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i] < buy {
			buy = prices[i]
			continue
		}

		profit := prices[i] - buy
		if profit > ans {
			ans = profit
		}
	}

	return ans
}
