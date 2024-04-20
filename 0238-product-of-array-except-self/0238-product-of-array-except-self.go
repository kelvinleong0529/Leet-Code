func productExceptSelf(nums []int) []int {
	ans := make([]int, len(nums))
	ans[0] = 1

	// forward iteration
	for i := 0; i < len(nums)-1; i++ {
		ans[i+1] = ans[i] * nums[i]
	}

	// backward iteration
	backwardMultiplier := nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		ans[i] *= backwardMultiplier
		backwardMultiplier *= nums[i]
	}

	return ans
}