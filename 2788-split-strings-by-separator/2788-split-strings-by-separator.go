func splitWordsBySeparator(words []string, separator byte) []string {
	answer := make([]string, 0)
	for i := range words {
		var word string
		for _, v := range words[i] {
			if v == rune(separator) {
				if word != "" {
					answer = append(answer, word)
					word = ""
				}
			} else {
				word += string(v)
			}
		}
		if word != "" {
			answer = append(answer, word)
		}
	}
	return answer
}
