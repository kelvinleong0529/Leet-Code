func getCommon(nums1 []int, nums2 []int) int {
	index1 := 0
	index2 := 0
	for {
		if index1 >= len(nums1) || index2 >= len(nums2) {
			return -1
		}

		num1 := nums1[index1]
		num2 := nums2[index2]
		
		if num1 == num2 {
			return num1
		}
		
		if num1 > num2 {
			index2++
		} else {
			index1++
		}
	}
}