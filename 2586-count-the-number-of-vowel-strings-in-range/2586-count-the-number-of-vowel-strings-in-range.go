func vowelStrings(words []string, left int, right int) int {
	var answer int
	for i := left; i <= right; i++ {
		word := words[i]
		start := word[0]
		end := word[len(word)-1]
		if isVowel(start) && isVowel(end) {
			answer++
		}
	}
	return answer
}

func isVowel(char byte) bool {
	switch char {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	}
	return false
}