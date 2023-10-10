func uniqueOccurrences(arr []int) bool {
    hashMap := make(map[int]int)
	for _, v := range arr {
		hashMap[v]++
	}
	occurenceHashMap := make(map[int]int)
	for _, v := range hashMap {
		occurenceHashMap[v]++
	}
	for _, v := range occurenceHashMap {
		if v != 1 {
			return false
		}
	}
	return true
}