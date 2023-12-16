func maximumValue(strs []string) int {
	answer := 0
	for _, str := range strs {
		value := getLength(str)
		if value > answer {
			answer = value
		}
	}
	return answer
}

func getLength(str string) int {
	if isNumeric(str) {
		num, _ := strconv.ParseInt(str, 10, 64)
		return int(num)
	} else {
		return len(str)
	}
}

func isNumeric(str string) bool {
	for _, char := range str {
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}