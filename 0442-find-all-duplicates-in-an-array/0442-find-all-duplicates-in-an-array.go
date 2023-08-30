func findDuplicates(nums []int) []int {
    hashMap := make(map[int]bool)
	answer := make([]int, 0)
	for _, v := range nums {
		if hashMap[v] {
			answer = append(answer, v)
		} else {
			hashMap[v] = true
		}
	}
	return answer
}