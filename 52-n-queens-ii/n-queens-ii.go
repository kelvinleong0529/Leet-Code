func totalNQueens(n int) int {
    col := make(map[int]bool)
	posDiagonal := make(map[int]bool)
	negDiagonal := make(map[int]bool)
    ans := 0

	var dfs func(int)
	dfs = func(r int) {
		if r == n {
			ans++
			return
		}

		for c := range n {
			posD := r + c
			negD := r - c

			if !col[c] && !posDiagonal[posD] && !negDiagonal[negD] {
				col[c] = true
				posDiagonal[posD] = true
				negDiagonal[negD] = true

				dfs(r + 1)

				col[c] = false
				posDiagonal[posD] = false
				negDiagonal[negD] = false
			}
		}
	}

	dfs(0)

	return ans
}