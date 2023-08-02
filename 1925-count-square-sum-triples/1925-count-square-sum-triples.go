func countTriples(n int) int {
    squares := make([]int, 0)
    tripletsSum := make(map[int]int)
    var answer int
    for i := 1; i <= n; i ++ {
        square := i * i
        if v, ok := tripletsSum[square]; ok {
            answer += 2 * v
        }
        if len(squares) != 0 {
           for _, v := range squares {
               sum := square + v
               tripletsSum[sum]++
            } 
        }
        squares = append(squares, square)
    }
    return answer
}