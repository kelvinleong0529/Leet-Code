func occurrencesOfElement(nums []int, queries []int, x int) []int {
	indexList := make([]int, 0)
	for i, num := range nums {
		if num == x {
			indexList = append(indexList, i)
		}
	}

	ans := make([]int, len(queries))
	for i, query := range queries {
		var index int
		if query > len(indexList) {
			index = -1
		} else {
			index = indexList[query-1]
		}
		ans[i] = index
	}
	
	return ans
}