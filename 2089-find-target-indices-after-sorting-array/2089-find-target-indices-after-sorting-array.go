func targetIndices(nums []int, target int) []int {
	array := make([]int, 0)
	for _, num := range nums {
		for {
			if len(array) > num-1 {
				break
			}
			array = append(array, 0)
		}
		array[num-1]++
	}
	
	count := 0
	ans := make([]int, 0)
	for i := range array {
		if i == target -1 {
			for j := 0; j < array[i]; j++ {
				ans = append(ans, count + j)
			}
			break
		}
		count += array[i]
	}
	return ans
}