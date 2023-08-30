func shortestPalindrome(s string) string {
    if isPalindrome(s) {
        return s
    }
	reversed := reverseString(s)
	var count int
	for i := range s {
		word := s[:i]
		if word == reversed[len(reversed)-i:] {
			count = i
		}
	}
	return reversed[:len(reversed)-count] + s
}

func reverseString(s string) string {
	var reversed string
	for i := len(s) - 1; i >= 0; i-- {
		reversed += string(s[i])
	}
	return reversed
}

func isPalindrome(s string) bool {
    for i, j := 0, len(s)-1; i< j; i,j = i+1, j-1 {
        if s[i] != s[j] {
            return false
        }
    }
    return true
}