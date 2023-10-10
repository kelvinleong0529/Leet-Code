func countConsistentStrings(allowed string, words []string) int {
	hashMap := make(map[rune]bool)
	for _, v := range allowed {
		hashMap[v] = true
	}

	var answer int
	for _, v := range words {
		if validString(hashMap, v) {
			answer++
		}
	}
	return answer
}

func validString(hashMap map[rune]bool, word string) bool {
	for _, v := range word {
		if _, ok := hashMap[v]; !ok {
			return false
		}
	}
	return true
}