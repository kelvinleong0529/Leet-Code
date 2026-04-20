func convert(s string, numRows int) string {
	r := 0
	dir := true
	arr := make([]*strings.Builder, numRows)
	for i := range numRows {
		arr[i] = &strings.Builder{}
	}

	updateDir := func() {
		if dir {
			if r < numRows-1 {
				r++
			} else {
				r--
				dir = !dir
			}
		} else {
			if r > 0 {
				r--
			} else {
				r++
				dir = !dir
			}
		}
	}

	for _, v := range s {
		arr[r].WriteRune(v)
		updateDir()
	}

	ans := &strings.Builder{}
	for i := range numRows {
		ans.WriteString(arr[i].String())
	}

	return ans.String()
}

func intToRoman(num int) string {
	intArr := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	strArr := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	ans := strings.Builder{}
	i := 0

	for num > 0 {
		quotient := num / intArr[i]
		for range quotient {
			ans.WriteString(strArr[i])
		}
		num = num % intArr[i]
	}

	return ans.String()
}

func countAndSay(n int) string {

	var dfs func(int) string
	dfs = func(i int) string {
		if i == 1 {
			return "1"
		}

		prev := dfs(i - 1)
		str := strings.Builder{}
		count := 1
		var current rune

		for i, v := range prev {
			if i == 0 {
				current = v
				continue
			}
			if v == current {
				count++
				continue
			} else {
				str.WriteRune(rune(count + '0'))
				str.WriteRune(current)
				current = v
				count = 1
			}
		}

        str.WriteRune(rune(count + '0'))
		str.WriteRune(current)

		return str.String()
	}

	return dfs(n)
}
