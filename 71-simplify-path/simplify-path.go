func simplifyPath(path string) string {
	stack := make([]string, 0)
	parts := strings.Split(path, "/")

	for _, part := range parts {
		if len(part) == 0 {
			continue
		}
		if part == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
			continue
		}
		if part == "." {
			continue
		}
		stack = append(stack, part)
	}

	if len(stack) == 0 {
		return "/"
	}

	var ans strings.Builder

	for _, v := range stack {
		ans.WriteString("/")
		ans.WriteString(v)
	}

	return ans.String()
}