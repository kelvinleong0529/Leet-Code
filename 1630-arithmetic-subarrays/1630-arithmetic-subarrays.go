func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	answer := make([]bool, len(l))
	for i := range l {
		start := l[i]
		end := r[i]
		input := make([]int, end-start+1)
		copy(input, nums[start:end+1])
		answer[i] = checkIsArithmeticArray(input)
	}
	return answer
}

func checkIsArithmeticArray(nums []int) bool {
	if len(nums) < 2 {
		return false
	}
	if len(nums) == 2 {
		return true
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	difference := nums[1] - nums[0]
	for i := 1; i <= len(nums)-2; i++ {
		if nums[i+1]-nums[i] != difference {
			return false
		}
	}
	return true
}