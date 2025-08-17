func climbStairs(n int) int {
	plusTwo := 0
	plusOne := 0
	current := 1

	for i := n; i > 0; i-- {
		plusTwo = plusOne
		plusOne = current
		current = plusTwo + plusOne
	}

	return current
}