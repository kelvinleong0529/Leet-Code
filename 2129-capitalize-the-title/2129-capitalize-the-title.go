func capitalizeTitle(title string) string {
	var answer string
	word := make([]byte, 0)
	for _, v := range title {
		if v == ' ' {
			if len(word) > 2 {
				word[0] = byte(capitalize(rune(word[0])))
			}
			answer += string(word) + " "
			word = []byte{}
		} else {
			word = append(word, byte(lowercase(v)))
		}
	}
	if len(word) > 2 {
		word[0] = byte(capitalize(rune(word[0])))
	}
	return answer + string(word)
}

func capitalize(char rune) rune {
	if char >= 'a' && char <= 'z' {
		return char - 'a' + 'A'
	}
	return char
}

func lowercase(char rune) rune {
	if char >= 'A' && char <= 'Z' {
		return char + 'a' + -'A'
	}
	return char
}
