func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := maxCandies(candies)
	answer := make([]bool, len(candies))
	for i, v := range candies {
		answer[i] = v+extraCandies >= max
	}
	return answer
}

func maxCandies(candies []int) int {
	var answer int
	for _, v := range candies {
		if v > answer {
			answer = v
		}
	}
	return answer
}