func generate(numRows int) [][]int {
	dp := make([][]int, 0)

	for i := 0; i < numRows; i++ {
		row := make([]int, i+1)
		for col := 1; col < len(row)-1; col++ {
			row[col] = dp[i-1][col-1] + dp[i-1][col]
		}
		row[0] = 1
		row[len(row)-1] = 1
		dp = append(dp, row)
	}

	return dp[:numRows]
}