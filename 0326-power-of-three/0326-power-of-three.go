func isPowerOfThree(n int) bool {
	if n == 1 {
		return true
	}
	for {
		if n == 3 {
			return true
		}
		if n%3 != 0 || n < 3 {
			return false
		}
		n = n / 3
	}
}