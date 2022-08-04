func max(num1 int, num2 int) int {
	if num1 < num2 {
		return num2
	}
	return num1
}

func maxProfit(prices []int) int {
    maxProfit := 0
	leftPointer := 0
	rightPointer := 1

	for {

		if rightPointer == len(prices) {
			break
		}

		if prices[leftPointer] > prices[rightPointer] {
			leftPointer = rightPointer
			rightPointer++
			continue
		}

		profit := prices[rightPointer] - prices[leftPointer]
		maxProfit = max(profit, maxProfit)
		rightPointer++
	}

	return maxProfit
}