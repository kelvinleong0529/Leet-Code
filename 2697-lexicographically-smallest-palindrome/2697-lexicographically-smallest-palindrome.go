func makeSmallestPalindrome(s string) string {
	answer := []byte(s)
	for i, j := 0, len(answer)-1; i < j; i, j = i+1, j-1 {
		if answer[i] == answer[j] {
			continue
		}
		if answer[i] < answer[j] {
			answer[j] = answer[i]
		} else {
			answer[i] = answer[j]
		}
	}
	return string(answer)
}