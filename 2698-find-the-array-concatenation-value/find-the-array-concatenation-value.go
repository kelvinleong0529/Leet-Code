func findTheArrayConcVal(nums []int) int64 {
	var ans int64
	for len(nums) > 0 {
		if len(nums) == 1 {
			ans += int64(nums[0])
            break
		}
		str := strconv.Itoa(nums[0]) + strconv.Itoa(nums[len(nums)-1])
		val, _ := strconv.Atoi(str)
		ans += int64(val)
		nums = nums[1:len(nums)-1]
	}
	return ans
}