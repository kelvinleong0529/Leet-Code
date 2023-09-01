func isValid(s string) bool {
	stack := make([]rune, 0)
	hashMap := map[rune]rune{
		'}': '{',
		')': '(',
		']': '[',
	}
	for _, v := range s {
		switch v {
		case ')', ']', '}':
            if len(stack) == 0 {
				return false
			}
			pop := stack[len(stack)-1]
			if hashMap[v] != pop {
				return false
			}
			stack = stack[:len(stack)-1]
		default:
			stack = append(stack, v)
		}
	}
	return len(stack) == 0
}