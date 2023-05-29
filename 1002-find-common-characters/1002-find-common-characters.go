func commonChars(words []string) []string {
	hashMap := make(map[rune]map[int]int)

	for i := range words {
		for _, word := range words[i] {
			if _, ok := hashMap[word]; !ok {
				hashMap[word] = make(map[int]int)
			}
			if _, ok := hashMap[word][i]; !ok {
				hashMap[word][i] = 0
			}
			hashMap[word][i] += 1
		}
	}

	ans := make([]string, 0)

	for word, counts := range hashMap {
		maxCount := len(words)
		for _, count := range counts {
			if len(counts) != len(words) {
				maxCount = 0
				continue
			}
			if count < maxCount {
				maxCount = count
			}
		}
		for i := 0; i < maxCount; i++ {
			ans = append(ans, string(word))
		}
	}

	return ans
}