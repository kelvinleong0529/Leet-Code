func checkDistances(s string, distance []int) bool {
	distanceMap := make(map[int]int)
	occurs := make(map[rune]int)
	for i, v := range s {
		if index, ok := occurs[v]; ok {
			distanceMap[int(v)-int('a')] = i - index - 1
		} else {
			occurs[v] = i
		}
	}
	for i := range distance {
		v, ok := distanceMap[i]
		if ok {
			if v != distance[i] {
				return false
			}
		}
	}
	return true
}