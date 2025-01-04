func rob(nums []int) int {
	maxRobbed := 0
	robbed := make([]int, len(nums))
	for i, num := range nums {
		var currentRob int
		if i == 0 || i == 1 {
			currentRob = num
		} else {
			currentRob = num + robbed[i-2]
		}
		if currentRob > maxRobbed {
			maxRobbed = currentRob
		}
		robbed[i] = maxRobbed
	}
	return maxRobbed
}