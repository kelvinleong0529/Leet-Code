func executeInstructions(n int, startPos []int, s string) []int {
	answer := make([]int, len(s))

	numOfSteps := func(s string) int {
		x := startPos[0]
		y := startPos[1]
		var count int
		for _, v := range s {
			switch string(v) {
			case "U":
				x -= 1
			case "D":
				x += 1
			case "L":
				y -= 1
			case "R":
				y += 1
			}
			if x < 0 || x > n-1 || y < 0 || y > n-1 {
				break
			}
			count++
		}
		return count
	}

	for i := range s {
		answer[i] = numOfSteps(s[i:])
	}

	return answer
}