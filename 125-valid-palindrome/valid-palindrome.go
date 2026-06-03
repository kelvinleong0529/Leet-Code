func isPalindrome(s string) bool {
	l, r := 0, len(s)-1

	for l <= r {
		if !isAlphanumeric(s[l]) {
			l++
			continue
		}
		if !isAlphanumeric(s[r]) {
			r--
			continue
		}
		if toLowerCase(s[l]) != toLowerCase(s[r]) {
			return false
		}
		l++
		r--
	}

	return true
}

func isAlphanumeric(b byte) bool {
	return (b >= '0' && b <= '9') ||
		(b >= 'a' && b <= 'z') ||
		(b >= 'A' && b <= 'Z')
}

func toLowerCase(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b + 32
	}
	return b
}
