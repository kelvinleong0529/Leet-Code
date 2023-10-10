func sumOddLengthSubarrays(arr []int) int {
	var answer int
	length := 1
	for {
		if length > len(arr) {
			break
		}
		for i := range arr {
			if i+length > len(arr) {
				break
			}
			answer += sumArray(arr[i : i+length])
		}
		length += 2
	}
	return answer
}

func sumArray(arr []int) int {
	var answer int
	for _, v := range arr {
		answer += v
	}
	return answer
}