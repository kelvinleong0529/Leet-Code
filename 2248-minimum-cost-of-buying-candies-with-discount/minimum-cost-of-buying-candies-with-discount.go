func minimumCost(cost []int) int {
	sort.Ints(cost)

	ans := 0
	index := len(cost) - 1

	for {
		if index <= 1 {
			for i := index; i >=0; i-- {
				ans += cost[i]
			}
			break
		}
		ans += cost[index] + cost[index-1]
		index -= 3
	}

	return ans
}