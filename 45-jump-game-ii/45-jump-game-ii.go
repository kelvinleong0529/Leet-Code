func jump(nums []int) int {
	da := make(map[int]int)
	da[0] = 0
	for i := 1; i < len(nums); i++ {
		if nums[len(nums)-i-1] == 0 {
			continue
		}
		temp := i - nums[len(nums)-i-1]
		if temp <= 0 {
			da[i] = 1
			continue
		}
		min := len(nums)
		for j := temp; j < i; j++ {
			v, ok := da[j]
			if ok {
				newValue := 1 + v
				if newValue < min {
					min = newValue
				}
			}
		}
		da[i] = min
	}
	return da[len(nums)-1]
}