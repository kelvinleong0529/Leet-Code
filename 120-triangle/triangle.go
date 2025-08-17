func minimumTotal(triangle [][]int) int {
	for row := 1; row < len(triangle); row++ {
		if row == 1 {
			triangle[row][0] += triangle[0][0]
			triangle[row][1] += triangle[0][0]
			continue
		}
		for col := 0; col < len(triangle[row]); col++ {
			if col == 0 {
				triangle[row][col] += triangle[row-1][col]
			} else if col == len(triangle[row])-1 {
				triangle[row][col] += triangle[row-1][len(triangle[row-1])-1]
			} else {
				smaller := min(triangle[row-1][col-1], triangle[row-1][col])
				triangle[row][col] += smaller
			}
		}
	}

	return minSlice(triangle[len(triangle)-1])
}

func minSlice(nums []int) int {
	ans := 10_001
	for _, num := range nums {
		if num < ans {
			ans = num
		}
	}
	return ans
}