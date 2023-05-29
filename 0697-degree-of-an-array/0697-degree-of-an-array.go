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
		if v[0] > maxFrequency {
			maxFrequency = v[0]
			distance := v[2] - v[1]
			if distance == 0 {
				continue
			}
			ans = distance
		} else if v[0] == maxFrequency {
			distance := v[2] - v[1]
			if distance == 0 {
				continue
			}
			if distance < ans {
				ans = distance
			}
		}
	}
    if maxFrequency == 1 {
		return 1
	}
	return ans + 1
}