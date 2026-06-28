func isValid(s string) bool {
	stack := make([]rune, 0)
	hashMap := map[rune]rune{')': '(', '}': '{', ']': '['}

	peek := func() (rune, bool) {
		if len(stack) == 0 {
			return 0, false
		}
		return stack[len(stack)-1], true
	}

	pop := func() {
		stack = stack[:len(stack)-1]
	}

	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
			continue
		}
		expected := hashMap[char]
		top, ok := peek()
		if !ok || top != expected {
			return false
		}
		pop()
	}

	return len(stack) == 0
}