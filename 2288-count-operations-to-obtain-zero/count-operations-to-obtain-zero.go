func countOperations(num1 int, num2 int) int {
	count := 0
	for {
        if num1 == 0 || num2 == 0 {
            return count
        }
		if num1 >= num2 {
			if num1%num2 == 0 {
				return count + num1/num2
			}
			num1 -= num2
			count++
		} else {
			if num2%num1 == 0 {
				return count + num2/num1
			}
			num2 -= num1
			count++
		}
	}
}