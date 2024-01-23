func findDuplicate(nums []int) int {
	array := make([]int, len(nums))
	for _, num := range nums {
        if array[num-1] == -1 {
			return num
		}
		array[num-1]--
	}
	return 0
}