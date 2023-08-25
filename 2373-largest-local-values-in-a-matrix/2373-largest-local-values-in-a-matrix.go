func largest3x3(grid [][]int, r1, r2, c1, c2 int) int {
	answer := grid[r1][c1]
	for i := r1; i < r2; i++ {
		for j := c1; j < c2; j++ {
			if grid[i][j] > answer {
				answer = grid[i][j]
			}
		}
	}
	return answer
}

func largestLocal(grid [][]int) [][]int {
	row := len(grid) - 2
	col := len(grid[0]) - 2
	answer := make([][]int, row)
	for i := 0; i < row; i++ {
		answer[i] = make([]int, col)
		for j := 0; j < col; j++ {
			answer[i][j] = largest3x3(grid, i, i+3, j, j+3)
		}
	}
	return answer
}