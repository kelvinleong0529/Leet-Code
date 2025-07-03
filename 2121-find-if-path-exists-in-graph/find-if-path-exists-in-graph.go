func validPath(n int, edges [][]int, source int, destination int) bool {

    if source == destination {
        return true
    }

	hashMap := make(map[int][]int)

	for _, edge := range edges {
		edgeSource := edge[0]
		edgeDestination := edge[1]
		if _, ok := hashMap[edgeSource]; !ok {
			hashMap[edgeSource] = make([]int, 0)
		}
		hashMap[edgeSource] = append(hashMap[edgeSource], edgeDestination)

		if _, ok := hashMap[edgeDestination]; !ok {
			hashMap[edgeDestination] = make([]int, 0)
		}
		hashMap[edgeDestination] = append(hashMap[edgeDestination], edgeSource)
	}

	visitedNode := make(map[int]bool)
	queue := []int{source}

	for {
		if len(queue) == 0 {
			return false
		}

		currentNode := queue[0]
		queue = queue[1:]

		for _, nodeDestination := range hashMap[currentNode] {
			if nodeDestination == destination {
				return true
			}
			if !visitedNode[nodeDestination] {
				queue = append(queue, nodeDestination)
				visitedNode[nodeDestination] = true
			}
		}
	}
}