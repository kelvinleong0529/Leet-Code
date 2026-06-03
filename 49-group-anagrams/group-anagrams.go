func groupAnagrams(strs []string) [][]string {
	
	hashMap := make(map[string][]string)
	for _, str := range strs {
		b := []byte(str)
		sort.Slice(b, func(i int, j int) bool {return b[i] < b[j]})
		s := string(b)
		hashMap[s] = append(hashMap[s], str)
	}

	ans := make([][]string, 0)
	for _, v := range hashMap {
		ans = append(ans, v)
	}

	return ans
}