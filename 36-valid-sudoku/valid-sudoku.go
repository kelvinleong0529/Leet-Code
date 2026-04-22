func isValidSudoku(board [][]byte) bool {
	row := make([]int, 9)
	col := make([]int, 9)
	square := make([][]int, 3)

	for i := range len(square) {
		square[i] = make([]int, 3)
	}

	for r := range len(board) {
		for c := range len(board[r]) {
			if board[r][c] == '.' {
				continue
			}
			num := board[r][c] - '0' - 1
			squareRow := r / 3
			squareCol := c / 3
			if row[r]&(1<<num) != 0 ||
				col[c]&(1<<num) != 0 ||
				square[squareRow][squareCol]&(1<<num) != 0 {
				return false
			}
			row[r] |= (1 << num)
			col[c] |= (1 << num)
			square[squareRow][squareCol] |= (1 << num)
		}
	}

	return true
}