func reverseOnlyLetters(s string) string {
	index := make([]int, 0)
	runes := make([]rune, 0)
	for i, v := range s {
		if (v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z') {
			index = append(index, i)
			runes = append(runes, v)
		}
	}
	reversedRunes := reverse(runes)
	answer := []byte(s)
	for i, v := range index {
		answer[v] = byte(reversedRunes[i])
	}
	return string(answer)
}

func reverse[T any](slice []T) []T {
	answer := make([]T, len(slice))
	for i, j := len(slice)-1, 0; i >= 0; i, j = i-1, j+1 {
		answer[j] = slice[i]
	}
	return answer
}