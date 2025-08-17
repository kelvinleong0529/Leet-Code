func uniquePathsWithObstacles(obstacleGrid [][]int) int {

	rowLen := len(obstacleGrid)
	columnLen := len(obstacleGrid[0])

	grid := make([][]int, rowLen+1)
	for row := range grid {
		grid[row] = make([]int, columnLen+1)
	}

	for row := rowLen - 1; row >= 0; row-- {
		for col := columnLen - 1; col >= 0; col-- {
			if obstacleGrid[row][col] == 1 {
				grid[row][col] = 0
			} else if row == rowLen-1 && col == columnLen-1 {
				grid[row][col] = 1
			} else {
				grid[row][col] = grid[row+1][col] + grid[row][col+1]
			}
		}
	}

	return grid[0][0]
}