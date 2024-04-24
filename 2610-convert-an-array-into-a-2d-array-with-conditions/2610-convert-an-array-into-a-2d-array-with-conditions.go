func findMatrix(nums []int) [][]int {
	answer := make([][]int, 0)
	hashMap := make(map[int]int)
	
	for _, v := range nums {
		hashMap[v]++;
	}
	
	for {
		if len(hashMap) == 0{
			return answer
		}
		arr := make([]int, 0)
		for k, v := range hashMap {
			arr = append(arr, k)
			value := v-1
			if value == 0 {
				delete(hashMap, k)
			} else {
				hashMap[k] = value
			}
		}
		answer = append(answer, arr)
	}

}