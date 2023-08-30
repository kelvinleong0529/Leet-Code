func removeTrailingZeros(num string) string {
	answer := num
	var remove int
	for i := len(num) - 1; i >= 0; i-- {
		if num[i] == '0' {
			remove++
		} else {
			break
		}
	}

	return answer[:len(answer)-remove]
}