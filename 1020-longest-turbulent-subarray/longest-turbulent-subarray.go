type sign int

const (
	smaller sign = iota
	bigger
	equal
)

func maxTurbulenceSize(arr []int) int {
	ans := 1
	l := 0
	for r := 1; r < len(arr); r++ {
		currentSign := getSign(arr[r-1], arr[r])
		if currentSign == equal {
			l = r
		} else if r == l+1 {
			// window has length 2, just extend
		} else {
			prevSign := getSign(arr[r-2], arr[r-1])
			if !isOpposite(prevSign, currentSign) {
				l = r - 1
			}
		}
		ans = max(ans, r-l+1)
	}
	return ans
}

func isOpposite(sign1, sign2 sign) bool {
	if sign1 == smaller && sign2 == bigger {
		return true
	}
	if sign1 == bigger && sign2 == smaller {
		return true
	}
	return false
}

func getSign(num1, num2 int) sign {
	switch {
	case num1 < num2:
		return smaller
	case num1 > num2:
		return bigger
	default:
		return equal
	}
}
