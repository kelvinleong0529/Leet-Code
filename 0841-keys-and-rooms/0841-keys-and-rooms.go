func canVisitAllRooms(rooms [][]int) bool {
	visited := make(map[int]bool)
	var dfs func(int)
	dfs = func(index int) {
		visited[index] = true
		for _, v := range rooms[index] {
			if _, ok := visited[v]; ok {
				continue
			}
			dfs(v)
		}
	}

	dfs(0)
	return len(visited) == len(rooms)
}