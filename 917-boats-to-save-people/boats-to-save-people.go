func numRescueBoats(people []int, limit int) int {
	ans := 0

	slices.Sort(people)

	l, r := 0, len(people)-1

	for l <= r {
        if l == r {
			ans++
			break
		}
		diff := limit - people[r]
		if diff >= people[l] {
			l++
		}
		r--
		ans++
	}

	return ans
}