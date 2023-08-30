func partitionLabels(s string) []int {
	hashMap := make(map[rune]int)
	answer := make([]int, 0)
	for i, v := range s {
		hashMap[v] = i
	}
	size := 0
	end := 0
	for i, v := range s {
		size++
		if hashMap[v] > end {
			end = hashMap[v]
		}
		if i == end {
			answer = append(answer, size)
			size = 0
		}
	}
	return answer
}