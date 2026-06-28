func solve(board [][]byte) {
	safe := make(map[int]bool)
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	var bfs func(int, int)
	bfs = func(r, c int) {
		queue := make([][]int, 0)
		queue = append(queue, []int{r, c})
		safe[computeIndex(r, c, len(board[0]))] = true
		for len(queue) > 0 {
			current := queue[0]
			queue = queue[1:]
			for _, direction := range directions {
				newRow := current[0] + direction[0]
				newCol := current[1] + direction[1]
				index := computeIndex(newRow, newCol, len(board[0]))
				if newRow >= 0 && newRow < len(board) &&
					newCol >= 0 && newCol < len(board[0]) &&
					board[newRow][newCol] == 'O' && !safe[index] {
					queue = append(queue, []int{newRow, newCol})
					safe[index] = true
				}
			}
		}
	}

	rows := []int{0, len(board) - 1}
	cols := []int{0, len(board[0]) - 1}

	for _, r := range rows {
		for c := range board[r] {
			index := computeIndex(r, c, len(board[0]))
			if board[r][c] == 'O' && !safe[index] {
				bfs(r, c)
			}
		}
	}

	for _, c := range cols {
		for r := range board {
			index := computeIndex(r, c, len(board[0]))
			if board[r][c] == 'O' && !safe[index] {
				bfs(r, c)
			}
		}
	}

	for r := range board {
		for c := range board[r] {
			index := computeIndex(r, c, len(board[0]))
			if _, ok := safe[index]; !ok && board[r][c] == 'O' {
				board[r][c] = 'X'
			}
		}
	}
}

func computeIndex(r, c, rowLength int) int {
	return r*rowLength + c
}
