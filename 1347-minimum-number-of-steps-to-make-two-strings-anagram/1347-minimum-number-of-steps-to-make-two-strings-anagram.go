func minSteps(s string, t string) int {
	hashMap1 := buildHashMap(s)
	hashMap2 := buildHashMap(t)
	visited := make(map[rune]bool)
	var answer int
	iterate := func(hashMap map[rune]int) {
		for k := range hashMap {
			if _, ok := visited[k]; ok {
				continue
			}
			visited[k] = true
			difference := hashMap1[k] - hashMap2[k]
			if difference < 0 {
				difference *= -1
			}
			answer += difference
		}
	}
	iterate(hashMap1)
	iterate(hashMap2)
	return answer / 2
}

func buildHashMap(input string) map[rune]int {
	answer := make(map[rune]int)
	for _, v := range input {
		answer[v]++
	}
	return answer
}