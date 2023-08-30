func isPrefixOfWord(sentence string, searchWord string) int {
	var word string
	var index int
	for _, v := range sentence {
		if v == ' ' {
			index++
			if isPrefix(searchWord, word) {
				return index
			}
			word = ""
		} else {
			word += string(v)
		}
	}
    index++
	if isPrefix(searchWord, word) {
		return index
	}
	return -1
}

func isPrefix(s1, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}
	return s1 == s2[:len(s1)]
}