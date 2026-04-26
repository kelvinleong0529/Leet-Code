func maximalSquare(matrix [][]byte) int {
	dp := make(map[int]int)
	max := -1

	rows, cols := len(matrix), len(matrix[0])
	var helper func(int, int) int
	helper = func(r, c int) int {
		if r >= rows || c >= cols {
			return 0
		}

		index := r*cols + c
		if _, ok := dp[index]; !ok {
			down := helper(r+1, c)
			right := helper(r, c+1)
			diag := helper(r+1, c+1)

			length := 0
			if matrix[r][c] == '1' {
				length = 1 + min(down, right, diag)
			}
			dp[index] = length
			if length > max {
				max = length
			}

		}

		return dp[index]
	}
	helper(0, 0)

	return max * max
}
