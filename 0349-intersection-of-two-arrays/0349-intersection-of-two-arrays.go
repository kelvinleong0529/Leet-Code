func intersection(nums1 []int, nums2 []int) []int {
	hashMap := make (map[int]bool)
	for _, num := range nums1 {
		hashMap[num] = false
	}
	for _, num := range nums2 {
		if _, ok := hashMap[num]; ok {
			hashMap[num] = true
		}
	}
	ans := make([]int, 0)
	for k, v := range hashMap {
		if v {
			ans = append(ans, k)
		}
	}
	return ans
}