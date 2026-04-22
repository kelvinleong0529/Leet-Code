func uniquePaths(m int, n int) int {
	row := make([]int, n)
	for i := range len(row) {
		row[i] = 1
	}

	for r := m - 2; r >= 0; r-- {
		for c := n - 2; c >= 0; c-- {
			row[c] = row[c] + row[c+1]
		}
	}

	return row[0]
}