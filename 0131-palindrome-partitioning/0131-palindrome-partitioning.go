func partition(s string) [][]string {
	ans := make([][]string, 0)
	part := make([]string, 0)
	var dfs func(int)
	dfs = func(index int) {
		if index >= len(s) {
            temp := make([]string, len(part))
            copy(temp, part)
			ans = append(ans, temp)
			return
		}
		for i := index; i < len(s); i++ {
			if isPalindrome(s[index:i+1]) {
				part = append(part, s[index:i+1])
				dfs(i + 1)
				part = part[:len(part)-1]
			}
		}
	}
	dfs(0)
	return ans
}

func isPalindrome(s string) bool {
    if len(s) == 0 {
        return false
    }
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}