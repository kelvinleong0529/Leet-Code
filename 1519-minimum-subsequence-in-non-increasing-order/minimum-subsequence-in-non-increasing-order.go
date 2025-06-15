func minSubsequence(nums []int) []int {
	sort.Ints(nums)

	totalSum := 0
	for i := range nums {
		totalSum += nums[i]
	}

	ans := make([]int, 0)
	ansSum := 0
	index := len(nums) - 1
	for i := index; i >= 0; i-- {
		ans = append(ans, nums[i])
		ansSum += nums[i]
		totalSum -= nums[i]
		if ansSum > totalSum {
			break
		}
	}

	return ans

}