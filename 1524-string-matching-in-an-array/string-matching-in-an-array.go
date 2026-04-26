func stringMatching(words []string) []string {
	ans := make([]string, 0)

	check := func(index int) bool {
		for i := 0; i < len(words); i++ {
			if i == index {
				continue
			}
			if len(words[i]) < len(words[index]) {
				continue
			}
			if strings.Contains(words[i], words[index]) {
				return true
			}
		}
		return false
	}

	for i := range len(words) {
		if check(i) {
			ans = append(ans, words[i])
		}
	}

	return ans
}