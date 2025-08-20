func primeSubOperation(nums []int) bool {
	primes := make(map[int]bool)

	previous := 0
	for _, num := range nums {
		start := num - previous - 1
		if start < 0 {
			return false
		}
		deductedPrime := false
		for i := start; i >= 2; i-- {
			_, ok := primes[i]
			if ok || isPrime(i) {
				primes[i] = true
				deductedPrime = true
				previous = num - i
				break
			}
		}
		if !deductedPrime {
			previous = num
		}
	}
	return true
}

func isPrime(n int) bool {
	upperBoundary := math.Sqrt(float64(n))
	for i := 2; i <= int(upperBoundary); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}