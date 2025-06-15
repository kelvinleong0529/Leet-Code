func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	answer := 0
	for i, v := range nums {
		if i%2 == 0 {
			answer += v
		}
	}
	return answer
}