func canCompleteCircuit(gas []int, cost []int) int {
	ans := -1

	diff := make([]int, len(gas))
	sum := 0
	for i := range gas {
		diff[i] = gas[i] - cost[i]
		sum += diff[i]
	}

	if sum < 0 {
		return ans
	}

	for i := range diff {
		valid := true
		currentSum := 0
		for j := i; j < len(diff); j++ {
			currentSum += diff[j]
			if currentSum < 0 {
				valid = false
				break
			}
		}
		if valid {
			return i
		}
	}

	return ans
}
