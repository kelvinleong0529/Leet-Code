func removeInvalidParentheses(s string) []string {
	ans := make([]string, 0)
	found := false
	visited := make(map[string]bool)
	queue := make([]string, 0)

	queue = append(queue, s)
	visited[s] = true

	for len(queue) > 0 && !found {
		currentBatchSize := len(queue)
		for range currentBatchSize {
			current := queue[0]
			queue = queue[1:]
			if isValidParentheses(current) {
				found = true
				ans = append(ans, current)
			}
			if found {
				continue
			}
			for i := range current {
				newString := removeCharByIndex(current, i)
				if !visited[newString] {
					queue = append(queue, newString)
					visited[newString] = true
				}
			}
		}
	}

	return ans
}

func isValidParentheses(s string) bool {
	count := 0
	for _, char := range s {
		if count < 0 {
			return false
		}
		switch char {
		case '(':
			count++
		case ')':
			count--
		}
	}

	return count == 0
}

func removeCharByIndex(s string, i int) string {
	if i < 0 || i >= len(s) {
		return ""
	}
	runes := []rune(s)
	return string(append(runes[:i], runes[i+1:]...))
}
