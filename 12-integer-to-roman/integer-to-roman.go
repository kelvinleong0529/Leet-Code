func intToRoman(num int) string {
    intArr := []int {1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
    strArr := []string {"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	ans := strings.Builder{}
	i := 0

	for num > 0 {
		quotient := num / intArr[i]
		for range quotient {
			ans.WriteString(strArr[i])
		}
		num = num % intArr[i]
        i++
	}

	return ans.String()
}