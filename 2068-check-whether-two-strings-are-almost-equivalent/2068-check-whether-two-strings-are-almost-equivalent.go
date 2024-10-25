func checkAlmostEquivalent(word1 string, word2 string) bool {
	hashMap1 := countOccurence(word1)
	hashMap2 := countOccurence(word2)

	for k, v := range hashMap1 {
		diff := v - hashMap2[k]
		if diff > 3 || diff < -3 {
			return false
		}
		
		delete(hashMap1, k)
		delete(hashMap2, k)
	}
	
	for _, v := range hashMap2 {
		if v > 3 {
			return false
		} 
	}
	
	return true
}

func countOccurence(word string) map[rune]int {
	hashMap := make(map[int32]int)
	for _, char := range word {
		hashMap[char]++
	}
	return hashMap
}