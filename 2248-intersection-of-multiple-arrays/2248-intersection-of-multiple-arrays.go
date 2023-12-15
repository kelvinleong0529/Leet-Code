func intersection(nums [][]int) []int {
	hashMap := make(map[int]int)
	for _, num := range nums {
		for _, element := range num {
			hashMap[element]++
		}
	}
	answer := make([]int, 0)
	for k, v := range hashMap {
		if v == len(nums) {
			answer = append(answer, k)
		}
	}
	sort.Ints(answer)
	return answer
}
