func xorOperation(n int, start int) int {
    result := 0

    for i := 0; i < n; i++ {
        num := start + 2*i
        result ^= num
    }

    return result
}