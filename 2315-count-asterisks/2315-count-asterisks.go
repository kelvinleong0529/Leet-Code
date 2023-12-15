func countAsterisks(s string) int {
	pair := false
	var count int;
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '|':
			pair = !pair
		case '*':
			if !pair {
				count++
			}
		}
	}
	return count
}