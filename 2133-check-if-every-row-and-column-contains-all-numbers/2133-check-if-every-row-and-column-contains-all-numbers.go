func transpose(arr [][]int) [][]int {
	var res = make([][]int, len(arr))
	for i := 0; i < len(arr); i++ {
		res[i] = make([]int, len(arr))
		for j := 0; j < len(arr); j++ {
			res[i][j] = arr[j][i]
		}
	}
	return res
}

func check(row []int) bool {
	hashMap := make(map[int]bool)
	for i := range row {
		_, ok := hashMap[row[i]]
		if ok {
			return false
		}
		hashMap[row[i]] = true
	}
	return true
}

func checkValid(matrix [][]int) bool {
	for i := range matrix {
		if !check(matrix[i]) {
			return false
		}
	}
	matrix = transpose(matrix)
	for i := range matrix {
		if !check(matrix[i]) {
			return false
		}
	}
	return true
}
