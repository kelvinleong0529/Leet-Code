func uniquePaths(m int, n int) int {
	grid := make([][]int, m+1)
	for row := range grid {
		grid[row] = make([]int, n+1)
	}

	for row := m - 1; row >= 0; row-- {
		for col := n - 1; col >= 0; col-- {
			if row == m-1 && col == n-1 {
				grid[row][col] = 1
				continue
			}

			grid[row][col] = grid[row+1][col] + grid[row][col+1]
		}
	}

	return grid[0][0]
}