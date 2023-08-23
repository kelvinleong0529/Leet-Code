func countPairs(nums []int, target int) int {
    var answer int
	hashMap := make(map[int]int)
	for _, num := range nums {
		for k ,v := range hashMap {
			if num + k < target {
				answer += v
			}
		}
		hashMap[num]++
	}
	return answer
}
