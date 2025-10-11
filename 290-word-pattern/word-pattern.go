func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")

	if len(pattern) != len(words) {
		return false
	}

	exist := make(map[string]bool)
	hashMap := make(map[rune]string)

	for i, ch := range pattern {
		word, ok := hashMap[ch]
		if ok {
			if word != words[i] {
				return false
			}
		} else {
			if exist[words[i]] {
				return false
			}
			hashMap[ch] = words[i]
			exist[words[i]] = true
		}
	}

	return true
}