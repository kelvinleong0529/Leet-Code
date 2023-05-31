func twoSum(numbers []int, target int) []int {
    leftPointer := 0
	rightPointer := len(numbers) - 1

	for {
		temp := numbers[leftPointer] + numbers[rightPointer]

		if temp == target {
			return []int{leftPointer + 1, rightPointer + 1}
		}

		if temp > target {
			rightPointer -= 1
		} else {
			leftPointer += 1
		}
	}
}