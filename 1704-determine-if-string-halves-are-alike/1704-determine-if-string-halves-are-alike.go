func halvesAreAlike(s string) bool {
	return numberOfVowels(s[:len(s)/2]) == numberOfVowels(s[len(s)/2:])
}

func numberOfVowels(s string) int {
	var count int
	for i := range s {
		switch s[i] {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			count++
		}
	}
	return count
}