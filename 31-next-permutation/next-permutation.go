
func nextPermutation(nums []int) {
	cur := -1
	pivot1 := -1

	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < cur {
			pivot1 = i
			break
		}

		cur = nums[i]
	}

	if pivot1 == -1 {
		reverse(nums, 0, len(nums)-1)
		return
	}

	pivot2 := -1

	for i := len(nums) - 1; i > pivot1; i-- {
		if nums[i] > nums[pivot1] {
			pivot2 = i
			break
		}
	}

    nums[pivot1], nums[pivot2] = nums[pivot2], nums[pivot1]

	reverse(nums, pivot1 + 1, len(nums)-1)
}

func reverse(nums []int, start int, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}