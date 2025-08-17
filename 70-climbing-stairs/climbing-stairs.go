func climbStairs(n int) int {
	ans := make([]int, n+2)
	ans[n] = 1

	for i := n; i > 0; i-- {
		if i == n {
			ans[i-1] = 1
		}
		ans[i-1] = ans[i] + ans[i+1]
	}

	return ans[0]
}