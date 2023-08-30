func longestNiceSubstring(s string) string {
	var dfs func(string) string
	dfs = func(s string) string {
		hashMap := make(map[rune]bool)
		for _, v := range s {
			hashMap[v] = true
		}
		if len(hashMap) < 2 {
			return ""
		}
		for i, v := range s {
			_, ok1 := hashMap[getLower(v)]
			_, ok2 := hashMap[getUpper(v)]
			if !ok1 || !ok2 {
				string1 := dfs(s[:i])
				string2 := dfs(s[i+1:])
                if len(string2) > len(string1) {
                    return string2
                } else {
                    return string1
                }
			}
		}
		return s
	}
	return dfs(s)
}

func getLower(char rune) rune {
	if char >= 'A' && char <= 'Z' {
		return char - 'A' + 'a'
	}
	return char
}

func getUpper(char rune) rune {
	if char >= 'a' && char <= 'z' {
		return char - 'a' + 'A'
	}
	return char
}