func cellsInRange(s string) []string {
	startCol := s[0]
	startRow := s[1]
	endCol := s[3]
	endRow := s[4]
	answer := make([]string, 0)
	for i := startCol; i <= endCol; i++ {
		for j := startRow; j <= endRow; j++ {
			answer = append(answer, string(i) + string(j))
		}
	}
	return answer
}