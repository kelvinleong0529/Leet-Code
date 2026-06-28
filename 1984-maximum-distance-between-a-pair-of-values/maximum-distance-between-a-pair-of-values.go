func maxDistance(nums1 []int, nums2 []int) int {
	ans := 0

	index1, index2 := 0, 0

	for {
		if index1 >= len(nums1) || index2 >= len(nums2) {
			break
		}
		if nums1[index1] <= nums2[index2] {
			tempIndex := index2
			for {
				if tempIndex >= len(nums2) {
					return ans
				}
				if nums1[index1] > nums2[tempIndex] {
					break
				}
				ans = max(tempIndex-index1, ans)
				tempIndex++
			}
		}

		index1++
		index2 = index1
	}

	return ans
}
