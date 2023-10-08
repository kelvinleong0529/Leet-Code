func mostWordsFound(sentences []string) int {
	var answer int
	for _, v := range sentences {
		numOfWords := countNumberOfWords(v)
		if numOfWords > answer {
			answer = numOfWords
		}
	}
	return answer
}

func countNumberOfWords(sentence string) int {
	if len(sentence) == 0 {
		return 0
	}
	answer := 1
	for _, v := range sentence {
		if v == ' ' {
			answer++
		}
	}
	return answer
}
