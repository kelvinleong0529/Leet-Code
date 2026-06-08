func timeRequiredToBuy(tickets []int, k int) int {
    ans := tickets[k]

	for i, v := range tickets {
		if i == k {
			continue
		} else if i < k {
			ans += min(v, tickets[k])
		} else {
			ans += min(v, tickets[k]-1)
		}
	}

	return ans
}