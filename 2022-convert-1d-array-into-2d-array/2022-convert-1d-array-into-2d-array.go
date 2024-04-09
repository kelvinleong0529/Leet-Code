func construct2DArray(original []int, m int, n int) [][]int {
	ans := make([][]int, m)
    if m*n != len(original) {
		return make([][]int, 0)
	}
	row := 0
	col := 0
	for _, v := range original {
		if col == 0 {
			ans[row] = make([]int, n)
		}
		ans[row][col] = v
		col++
		if col == n {
			row++
			col = 0
		}
	}
	return ans
}