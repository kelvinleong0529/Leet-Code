func canFinish(numCourses int, prerequisites [][]int) bool {
	prerequisitesMap := make(map[int][]int)
	prerequisitesChecked := make(map[int]bool)

	for _, v := range prerequisites {
		course := v[0]
		prerequisite := v[1]
		if _, ok := prerequisitesMap[course]; !ok {
			prerequisitesMap[course] = make([]int, 0)
		}
		prerequisitesMap[course] = append(prerequisitesMap[course], prerequisite)
		prerequisitesChecked[course] = false
	}

	var dfs func(int) bool
	visited := make(map[int]bool)
	dfs = func(course int) bool {
		if prerequisitesChecked[course] {
			return true
		}
		if visited[course] {
			return false
		}
		prerequisites := prerequisitesMap[course]
		if len(prerequisites) == 0 {
			return true
		}
		visited[course] = true
		for _, v := range prerequisites {
			if !dfs(v) {
				return false
			}
		}
		visited[course] = false
		prerequisitesChecked[course] = true
		return true
	}

	for course := range prerequisitesMap {
		if !dfs(course) {
			return false
		}
	}

	return true
}