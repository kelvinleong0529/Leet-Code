func solveNQueens(n int) [][]string {
	col := make(map[int]bool)
	posDiagonal := make(map[int]bool)
	negDiagonal := make(map[int]bool)

	ans := make([][]string, 0)
	curr := make([]string, 0)
	for range n {
		curr = append(curr, strings.Repeat(".", n))
	}

	var dfs func(int)
	dfs = func(r int) {
		if r == n {
			tmp := make([]string, n)
			copy(tmp, curr)
			ans = append(ans, tmp)
			return
		}

		for c := range n {
			posD := r + c
			negD := r - c

			if !col[c] && !posDiagonal[posD] && !negDiagonal[negD] {
				col[c] = true
				posDiagonal[posD] = true
				negDiagonal[negD] = true

				b := []byte(curr[r])
				b[c] = 'Q'
				curr[r] = string(b)

				dfs(r + 1)

				col[c] = false
				posDiagonal[posD] = false
				negDiagonal[negD] = false

				b = []byte(curr[r])
				b[c] = '.'
				curr[r] = string(b)
			}
		}
	}

	dfs(0)

	return ans
}