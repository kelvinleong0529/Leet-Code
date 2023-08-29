func sortVowels(s string) string {
    vowels := "aeiouAEIOU"
	vowelIndices := []int{}
	vowelChars := []byte{}

	// Identify vowel indices and collect vowel characters
	for i, char := range s {
		if strings.ContainsRune(vowels, char) {
			vowelIndices = append(vowelIndices, i)
			vowelChars = append(vowelChars, byte(char))
		}
	}

	// Sort the vowel characters in non-decreasing order
	sort.Slice(vowelChars, func(i, j int) bool {
		return vowelChars[i] < vowelChars[j]
	})

	// Create the resulting string t
	t := []byte(s)
	for i, index := range vowelIndices {
		t[index] = vowelChars[i]
    }

	return string(t)
}