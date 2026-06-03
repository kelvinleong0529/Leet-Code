func isAnagram(s string, t string) bool {
	hashMap := make(map[rune]int)

	for _, char := range s {
		hashMap[char]++
	}

	for _, char := range t {
		if _, ok := hashMap[char]; !ok {
			return false
		} else {
			hashMap[char]--
			if hashMap[char] == 0 {
				delete(hashMap, char)
			}
		}
	}

	return len(hashMap) == 0
}