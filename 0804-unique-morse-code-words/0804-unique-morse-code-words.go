var morse []string = []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

func uniqueMorseRepresentations(words []string) int {
	hashMap := make(map[string]bool)
	for _, v := range words {
		morseRepresentation := getMorseRepresentation(v)
		hashMap[morseRepresentation] = true
	}
	return len(hashMap)
}

func getMorseRepresentation(word string) string {
	var answer string
	for _, v := range word {
		asciiValue := charToAscii(v)
		answer += morse[asciiValue]
	}
	return answer
}

func charToAscii(char rune) int {
	return int(char - 'a')
}
