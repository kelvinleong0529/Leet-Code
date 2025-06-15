func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	
	result := 0
	gIndex := 0
	sIndex := 0
	
	for {
		if sIndex >= len(s) || gIndex >= len(g) {
			break
		}
		if s[sIndex] >= g[gIndex] {
			gIndex++
			result++
		}
		sIndex++
	}
	
	return result
}
