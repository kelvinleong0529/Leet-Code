func firstPalindrome(words []string) string {
	for _, v := range words {
		if isPalindrome(v) {
			return v
		}
	}
	return ""
}

func isPalindrome(input string) bool {
	for i := 0; i < len(input); i++ {
		if input[i] != input[len(input)-1-i] {
			return false
		}
	}
	return true
}
