func islandPerimeter(grid [][]int) int {
	ROW := len(grid)
	COL := len(grid[0])
	visited := make(map[int]map[int]bool)

	var dfs func(r, c int) int
	dfs = func(r, c int) int {
		value1, ok1 := visited[r]
		if ok1 {
			if _, ok2 := value1[c]; ok2 {
				return 0
			}
		}

		if r < 0 || c < 0 || r >= ROW || c >= COL || grid[r][c] == 0 {
			return 1
		}

		if !ok1 {
			visited[r] = make(map[int]bool)
		}
		visited[r][c] = true
		return dfs(r-1, c) + dfs(r+1, c) + dfs(r, c-1) + dfs(r, c+1)
	}

	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == 1 {
				return dfs(r, c)
			}
		}
	}
	return 0
}