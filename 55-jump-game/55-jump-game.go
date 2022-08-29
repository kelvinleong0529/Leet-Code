func canJump(nums []int) bool {
    if len(nums) == 1 {
        return true
    }
	total := len(nums)
	count := 0
	target := total - 2
	for i := target; i >= 0; i-- {
		count++
		if nums[i] >= count {
			target -= count
			count = 0
		}
	}
	if target < 0 {
		return true
	} else {
		return false
	}
}