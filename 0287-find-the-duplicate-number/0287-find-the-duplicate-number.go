func findDuplicate(nums []int) int {
	array := make([]int, len(nums))
	for _, num := range nums {
        println(num)
		array[num-1]--
		if array[num-1] == -2 {
			return num
		}
	}
	return 0
}