func isValidSudoku(board [][]byte) bool {
	
	row := make([][]int, 9)
	col := make([][]int, 9)
	square := make([][][]int, 3)

	for r := range len(board) {
		row[r] = make([]int, 9)
		col[r] = make([]int, 9)
		for c := range len(board[0]) {
			squareRow := r / 3
			squareCol := c / 3
            if square[squareRow] == nil {
                square[squareRow] = make([][]int, 3)
            }
			if square[squareRow][squareCol] == nil {
				square[squareRow][squareCol] = make([]int, 9)
			}
		}
	}

	for r := range len(board) {
		for c := range len(board[0]) {
			if board[r][c] == '.' {
				continue
			}
            num := board[r][c] - '0' - 1
			squareRow := r / 3
			squareCol := c / 3
			if row[r][num] != 0 || col[c][num] != 0 || square[squareRow][squareCol][num] != 0 {
				return false
			}
			row[r][num] = 1
			col[c][num] = 1
			square[squareRow][squareCol][num] = 1
		}
	}

	return true
}
