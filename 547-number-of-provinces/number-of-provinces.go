func findCircleNum(isConnected [][]int) int {
	ans := 0
	n := len(isConnected)
	visited := make(map[int]bool, 0)

	var dfs func(int)
	dfs = func(i int) {
		visited[i] = true
		for j := range n {
			if !visited[j] && isConnected[i][j] == 1 {
				dfs(j)
			}
		}
	}

	for i := range n {
		if visited[i] {
			continue
		}
		dfs(i)
		ans += 1
	}

	return ans
}