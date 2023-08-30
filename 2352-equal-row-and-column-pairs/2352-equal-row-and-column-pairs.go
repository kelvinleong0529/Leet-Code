func equalPairs(grid [][]int) int {
    n := len(grid)
	count := 0

	for ri := 0; ri < n; ri++ {
		for cj := 0; cj < n; cj++ {
			if compareRowsAndColumn(grid, ri, cj) {
				count++
			}
		}
	}

	return count
}

func compareRowsAndColumn(grid [][]int, ri, cj int) bool {
	n := len(grid)
	for i := 0; i < n; i++ {
		if grid[ri][i] != grid[i][cj] {
			return false
		}
	}
	return true
}