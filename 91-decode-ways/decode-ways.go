func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}

	dp := make([]int, len(s)+1)
	dp[len(dp)-1] = 1

	for i := len(s) - 1; i >= 0; i-- {
		if i == len(s)-1 {
			if isValid, _ := valid(s[i:]); isValid {
				dp[i] = 1
			}
			continue
		}
		if isValid, _ := valid(s[i : i+2]); isValid {
			dp[i] = dp[i+1] + dp[i+2]
		} else if isValid, _ := valid(s[i : i+1]); isValid {
			dp[i] = dp[i+1]
		}
	}

	return dp[0]
}


func valid(s string) (bool, error) {
    if s[0] == '0' {
		return false, nil
	}
	num, err := strconv.Atoi(s)
	if err != nil {
		return false, err
	}
	return num > 0 && num < 27, nil
}