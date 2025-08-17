func longestPalindrome(s string) string {
	var result string
	var resultLength int

	for i := range s {
		// odd case
		oddLengthResult, oddResultLength := findLongPalindrome(s, i, i)
		if oddResultLength > resultLength {
			resultLength = oddResultLength
            result = oddLengthResult
		}

		// even case
		evenLengthResult, evenResultLength := findLongPalindrome(s, i, i+1)
		if evenResultLength > resultLength {
            resultLength = evenResultLength
			result = evenLengthResult
		}
	}

	return result
}

func findLongPalindrome(s string, l int, r int) (string, int) {
	result := ""
	resultLength := 0
	for {
		if l < 0 || r > len(s)-1 || s[l] != s[r] {
			break
		}
		result = s[l : r+1]
		resultLength = len(result)
		l--
		r++
	}
	return result, resultLength
}