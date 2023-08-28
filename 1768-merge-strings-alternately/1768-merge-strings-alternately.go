func mergeAlternately(word1 string, word2 string) string {
	var answer string
	len1 := len(word1)
	len2 := len(word2)
	var longer int
	if len1 > len2 {
		longer = len1
	} else {
		longer = len2
	}

	for i := 0; i < longer; i++ {
		if i >= len(word1) {
			answer += word2[i : len(word2)]
			break
		}
		if i >= len(word2) {
			answer += word1[i : len(word1)]
			break
		}
		answer += string(word1[i]) + string(word2[i])
	}

	return answer
}