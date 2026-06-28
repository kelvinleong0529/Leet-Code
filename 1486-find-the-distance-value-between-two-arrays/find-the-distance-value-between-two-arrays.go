func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	ans := 0

	for _, v1 := range arr1 {
		valid := true
		for _, v2 := range arr2 {
			if getAbs(v1-v2) <= d {
				valid = false
				break
			}
		}
		if valid {
			ans++
		}
	}

	return ans
}

func getAbs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}