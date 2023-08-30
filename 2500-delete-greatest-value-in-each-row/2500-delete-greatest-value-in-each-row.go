func deleteGreatestValue(grid [][]int) int {
	var answer int
    num := len(grid[0])
	for i := 0; i < num; i++ {
		removed := make([]int, 0)
		for rowi, row := range grid {
			updated, element := removeGreatestElement(row)
			removed = append(removed, element)
			grid[rowi] = updated
		}
		index := greatestElement(removed)
		answer += removed[index]
	}
	return answer
}

func removeGreatestElement(slice []int) ([]int, int) {
	index := greatestElement(slice)
	greatestElement := slice[index]
	return removeElement(slice, index), greatestElement
}

func greatestElement(slice []int) int {
	var answer int
	var index int
	for i, v := range slice {
		if i == 0 {
			index = i
			answer = v
			continue
		}
		if v > answer {
			index = i
			answer = v
		}
	}
	return index
}

func removeElement(slice []int, index int) []int {
	answer := make([]int, len(slice)-1)
	copy(answer, append(slice[:index], slice[index+1:]...))
	return answer
}