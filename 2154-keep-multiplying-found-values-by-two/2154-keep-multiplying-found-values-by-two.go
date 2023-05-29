func findFinalValue(nums []int, original int) int {
	hashMap := make(map[int]bool)
	for i := range nums {
		hashMap[nums[i]] = true
	}
	for {
		if _, ok := hashMap[original]; ok {
			original *= 2
		} else {
			return original
		}
	}
}