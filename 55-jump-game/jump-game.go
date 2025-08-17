func canJump(nums []int) bool {
	dp := make([]bool, len(nums))
    dp[0] = true

	for i := range nums {
		for j := 0; j <= i; j++ {
			if dp[j] && j + nums[j] >= i {
				dp[i] = true
				break
			}
		}
	}
	
	return dp[len(nums)-1]
}