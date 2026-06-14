func findClosestElements(arr []int, k int, x int) []int {
	left, right := 0, len(arr)-1

	for right-left+1 > k {
		leftDiff := getAbs(x - arr[left])
		rightDiff := getAbs(arr[right] - x)

		if leftDiff <= rightDiff {
			right -= 1
		} else {
			left += 1
		}
	}

	return arr[left : right + 1]
}

func getAbs(num int) int {
	if num < 0 {
		return num * -1
	}
	return num
}