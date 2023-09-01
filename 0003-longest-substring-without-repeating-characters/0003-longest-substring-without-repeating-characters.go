func lengthOfLongestSubstring(s string) int {
    // Create a map to store the last seen index of each character
	charIndex := make(map[byte]int)
	
	maxLength := 0
	start := 0

	for end := 0; end < len(s); end++ {
		if lastIndex, exists := charIndex[s[end]]; exists {
			// If the character is repeated, update the start index to the next position
			start = max(start, lastIndex+1)
		}

		// Update the last seen index of the character
		charIndex[s[end]] = end

		// Calculate the length of the current substring
		currentLength := end - start + 1

		// Update the maximum length if needed
		maxLength = max(maxLength, currentLength)
	}

	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}