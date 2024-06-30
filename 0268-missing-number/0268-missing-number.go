func missingNumber(nums []int) int {
	expectedSum := len(nums) * (len(nums) + 1) / 2
	actualSum := 0
	for _, num := range nums {
		actualSum += num
	}
	return expectedSum - actualSum
}