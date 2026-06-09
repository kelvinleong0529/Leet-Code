func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{{}}
	}

	ans := make([][]int, 0)
	perms := permute(nums[1:])

	for _, v := range perms {
		for i := 0; i <= len(v); i++ {
			temp := make([]int, len(v))
			copy(temp, v)
			temp = slices.Insert(temp, i, nums[0])
			ans = append(ans, temp)
		}
	}

	return ans
}
