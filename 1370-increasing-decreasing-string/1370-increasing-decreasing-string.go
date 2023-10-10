func sortString(s string) string {
	hashMap := make(map[int]int)
	for _, v := range s {
		ascii := charToAscii(v)
		hashMap[ascii]++
	}

	var answer string
	for {
		if len(hashMap) == 0 {
			break
		}

		ascendingKeys := hashMapAscendingKeys(hashMap)
		for _, v := range ascendingKeys {
			answer += string(asciiToChar(v))
			hashMap[v]--
			if hashMap[v] == 0 {
				delete(hashMap, v)
			}
		}

		descendingKeys := hashMapDescendingKeys(hashMap)
		for _, v := range descendingKeys {
			answer += string(asciiToChar(v))
			hashMap[v]--
			if hashMap[v] == 0 {
				delete(hashMap, v)
			}
		}
	}

	return answer
}

func charToAscii(char rune) int {
	return int(char - 'a')
}

func asciiToChar(ascii int) rune {
	return rune(ascii + 'a')
}

func hashMapAscendingKeys(hashMap map[int]int) []int {
	answer := make([]int, 0)
	for k := range hashMap {
		answer = append(answer, k)
	}
	sort.Slice(answer, func(i, j int) bool {
		return answer[i] < answer[j]
	})
	return answer
}

func hashMapDescendingKeys(hashMap map[int]int) []int {
	answer := make([]int, 0)
	for k := range hashMap {
		answer = append(answer, k)
	}
	sort.Slice(answer, func(i, j int) bool {
		return answer[i] > answer[j]
	})
	return answer
}
