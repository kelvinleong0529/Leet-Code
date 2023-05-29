func findShortestSubArray(nums []int) int {
	hashMap := make(map[int][]int)
	for i, v := range nums {
		if _, ok := hashMap[v]; !ok {
			hashMap[v] = []int{0, i, i}
		}
		hashMap[v][0] += 1
		hashMap[v][2] = i
	}
	maxFrequency := 0
	ans := len(nums)
	for _, v := range hashMap {
		distance := v[2] - v[1]
		if distance == 0 {
			continue
		}
		if v[0] > maxFrequency {
			maxFrequency = v[0]
			ans = distance
		} else if v[0] == maxFrequency && distance < ans {
			ans = distance
		}
	}
	if maxFrequency == 0 {
		return 1
	}
	return ans + 1
}