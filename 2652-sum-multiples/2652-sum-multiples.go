func sumOfMultiples(n int) int {
    var answer int
    for i := 3; i <= n; i++ {
        if i % 3 == 0 || i % 5 == 0 || i % 7 == 0 {
            answer+=i
        }
    }
    return answer
}
