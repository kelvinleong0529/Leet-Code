func countGoodSubstrings(s string) int {
	var answer int
	for i := 0; i <= len(s)-3; i++ {
		if isGoodString(s[i : i+3]) {
			answer++
		}
	}
	return answer
}

func isGoodString(s string) bool {
	hashMap := make(map[rune]bool)
	for _, v := range s {
		if _, ok := hashMap[v]; ok {
			return false
		}
		hashMap[v] = true
	}
	return true
}