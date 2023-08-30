func digitCount(num string) bool {
	frequencies := make(map[int]int)
	expected := make(map[int]int)
	for i, v := range num {
		expected[i] = int(v - '0')
		frequencies[int(v-'0')]++
	}
	for k := range expected {
		if expected[k] != frequencies[k] {
			return false
		}
	}
	return true
}
