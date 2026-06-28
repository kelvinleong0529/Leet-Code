// guaranteed that c appears at least once in s
func shortestToChar(s string, c byte) []int {
	ans := make([]int, len(s))

	index := len(s)

	for i := 0; i < len(s); i++ {
		if s[i] == c {
			index = i
		}
		ans[i] = abs(i - index)
	}

	index = 0

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == c {
			index = i
		}
		ans[i] = min(ans[i], abs(i - index))
	}

	return ans
}

func abs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}