func updateMatrix(mat [][]int) [][]int {
	numOfRows := len(mat)
	numOfCols := len(mat[0])

	ans := make([][]int, numOfRows)
	queue := make([][]int, 0)

	for row := range numOfRows {
		ans[row] = make([]int, numOfCols)
		for col := range numOfCols {
			if mat[row][col] == 0 {
				ans[row][col] = 0
				queue = append(queue, []int{row, col})
			} else {
				ans[row][col] = -1
			}
		}
	}

	directions := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	for len(queue) > 0 {
		row, col := queue[0][0], queue[0][1]
		queue = queue[1:]
		for _, direction := range directions {
			newRow := row + direction[0]
			newCol := col + direction[1]

			if newRow < 0 || newRow >= numOfRows || newCol < 0 || newCol >= numOfCols {
				continue
			}

			newDistance := ans[row][col] + 1
			currentDistance := ans[newRow][newCol]
			if currentDistance == -1 || (currentDistance > 0 && newDistance < currentDistance) {
				ans[newRow][newCol] = newDistance
				queue = append(queue, []int{newRow, newCol})
			}
		}
	}

	return ans
}