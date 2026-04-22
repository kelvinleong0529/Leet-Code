func solveSudoku(board [][]byte) {
	var row, col, subgrid [9]int

	for r := range len(board) {
		for c := range len(board[r]) {
			if board[r][c] != '.' {
				mask := 1 << (board[r][c] - '1')
				row[r] |= mask
				col[c] |= mask
				subgrid[(r/3*3)+c/3] |= mask
			}
		}
	}

	var backtracking func(int, int) bool
	backtracking = func(r int, c int) bool {
		if c == 9 {
			return backtracking(r+1, 0)
		}
		if r == 9 {
			return true
		}
		if board[r][c] == '.' {
			for num := range 9 {
				mask := 1 << num
				subgridIndex := (r / 3 * 3) + c/3
				if row[r]&mask == 0 && col[c]&mask == 0 && subgrid[subgridIndex]&mask == 0 {
					board[r][c] = byte(num + '1')
					row[r] |= mask
					col[c] |= mask
					subgrid[subgridIndex] |= mask

					if backtracking(r, c+1) {
						return true
					}

					board[r][c] = '.'
					row[r] &= ^mask
					col[c] &= ^mask
					subgrid[subgridIndex] &= ^mask
				}
			}
            return false
		} else {
            return backtracking(r, c+1)
        }
	}

	backtracking(0, 0)
}