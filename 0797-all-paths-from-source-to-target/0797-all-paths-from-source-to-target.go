func allPathsSourceTarget(graph [][]int) [][]int {
	queue := make([][]int, 0)
	result := make([][]int, 0)
	target := len(graph) - 1

	queue = [][]int{{0}}

	for {
		if len(queue) == 0 {
			break
		}
		temp := queue[0]
		queue = queue[1:]
		if temp[len(temp)-1] == target {
			result = append(result, temp)
		} else {
			for _, neighbour := range graph[temp[len(temp)-1]] {
				path := make([]int, len(temp))
				copy(path, temp)
				queue = append(queue, append(path, neighbour))
			}
		}
	}

	return result
}