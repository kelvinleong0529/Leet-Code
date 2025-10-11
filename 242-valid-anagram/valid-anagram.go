func isAnagram(s string, t string) bool {
	hashMap := getHashMap(s)
	for _, ch := range t {
		count, ok := hashMap[ch]
		if !ok {
			return false
		}
		hashMap[ch] = count - 1
		if hashMap[ch] == 0 {
			delete(hashMap, ch)
		}
	}
	return len(hashMap) == 0
}

func getHashMap(s string) map[rune]int {
	hashMap := make(map[rune]int)
	for _, ch := range s {
		hashMap[ch]++
	}
	return hashMap
}
