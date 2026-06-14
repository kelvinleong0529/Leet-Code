func intersect(nums1 []int, nums2 []int) []int {
	ans := make([]int, 0)

	slices.Sort(nums1)
	slices.Sort(nums2)

	pointer1 := 0
	pointer2 := 0

	for {
		if pointer1 >= len(nums1) || pointer2 >= len(nums2) {
			break
		}

		if nums1[pointer1] == nums2[pointer2] {
			ans = append(ans, nums1[pointer1])
			pointer1 += 1
			pointer2 += 1
			continue
		}

		if nums1[pointer1] > nums2[pointer2] {
			pointer2 += 1
		} else {
			pointer1 += 1
		}
	}

	return ans
}
