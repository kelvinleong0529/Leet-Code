func minTimeToType(word string) int {
    var count int
	current := 'a'
	for _, v := range word {
        distance := int(v - current)
        if distance < 0 {
            distance = -distance
        }
        if 26-distance < distance {
            distance = 26-distance
        }
		count += distance + 1
		current = v
	}
	return count
}