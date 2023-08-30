func finalString(s string) string {
	var answer string
	for _, v := range s {
		if v == 'i' {
			answer = reverseString(answer)
		} else {
			answer += string(v)
		}
	}
	return answer
}

func reverseString(s string) string {
	var reversed string
	for i := len(s) - 1; i >= 0; i-- {
		reversed += string(s[i])
	}
	return reversed
}