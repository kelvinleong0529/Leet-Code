func minPathSum(grid [][]int) int {
	rowLen := len(grid)
	columLen := len(grid[0])

	bufferRow := make([]int, columLen+1)
	for i := 0; i < columLen+1; i++ {
		bufferRow[i] = 99999
	}
	bufferRow[len(bufferRow)-2] = 0

	for i := rowLen - 1; i >= 0; i-- {
		for j := columLen - 1; j >= 0; j-- {
			rightCost := grid[i][j] + bufferRow[j+1]
			downCost := grid[i][j] + bufferRow[j]

			cost := min(rightCost, downCost)
			bufferRow[j] = cost
		}
	}

	return bufferRow[0]
}