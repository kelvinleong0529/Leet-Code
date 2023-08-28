func minSteps(s string, t string) int {
	hashMap1 := buildHashMap(s)
	hashMap2 := buildHashMap(t)
	visited := make(map[rune]bool)
	var answer int
	for k := range hashMap1 {
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
	for k := range hashMap2 {
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
	return answer / 2
}

func buildHashMap(input string) map[rune]int {
	answer := make(map[rune]int)
	for _, v := range input {
		answer[v]++
	}
	return answer
}
