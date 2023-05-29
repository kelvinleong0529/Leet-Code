func exist(board [][]byte, word string) bool {

	ROW := len(board)
	COL := len(board[0])
	visited := make(map[int]map[int]bool)

	var dfs func(row, col, index int) bool
	dfs = func(row, col, index int) bool {
		if index == len(word) {
			return true
		}

		value1, ok1 := visited[row]
		if ok1 {
			value2, ok2 := value1[col]
			if ok2 && value2 {
				return false
			}
		}

		if row < 0 || col < 0 || row >= ROW || col >= COL || board[row][col] != word[index] {
			return false
		}

        if len(visited[row]) == 0 {
			visited[row] = make(map[int]bool)
		}
		visited[row][col] = true
		result := dfs(row-1, col, index+1) || dfs(row+1, col, index+1) || dfs(row, col-1, index+1) || dfs(row, col+1, index+1)
		delete(visited[row], col)
		if len(visited[row]) == 0 {
			delete(visited, row)
		}

		return result
	}

	for r := range board {
		for c := range board[r] {
			if dfs(r, c, 0) {
				return true
			}
		}
	}
	return false
}