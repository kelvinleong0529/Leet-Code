func sortTheStudents(score [][]int, k int) [][]int {
	result := make([][]int, len(score))

	exam := make([]int, len(score))
	examIndex := make(map[int]int)

	for i, v := range score {
		exam[i] = v[k]
		examIndex[v[k]] = i
	}

	sort.Slice(exam, func(i, j int) bool {
		return exam[i] > exam[j]
	})
    
	for i, v := range exam {
		index := examIndex[v]
		result[i] = score[index]
	}

	return result
}