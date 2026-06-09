func canFinish(numCourses int, prerequisites [][]int) bool {
	preMap := make(map[int][]int)

	for _, preprerequisite := range prerequisites {
		left := preprerequisite[0]
		right := preprerequisite[1]
		if _, ok := preMap[left]; !ok {
			preMap[left] = make([]int, 0)
		}
		preMap[left] = append(preMap[left], right)
	}

	visited := make(map[int]bool)

	var dfs func(int) bool

	dfs = func(n int) bool {
		if _, ok := visited[n]; ok {
			return false
		}
		if len(preMap[n]) == 0 {
			return true
		}

		visited[n] = true
		for _, v := range preMap[n] {
			if !dfs(v) {
				return false
			}
		}
		delete(visited, n)
		preMap[n] = []int{}

		return true
	}

	for i := 0; i < numCourses; i++ {
		if !dfs(i) {
			return false
		}
	}

	return true
}
