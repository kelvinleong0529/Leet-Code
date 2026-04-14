func zigzagTraversal(grid [][]int) []int {
    ans := make([]int, 0)

    forward := true
    skip := false

    for r := range len(grid) {
        if forward {
            for c := range len(grid[0]) {
                if !skip {
                    ans = append(ans, grid[r][c])
                }
                skip = !skip
            }
        } else {
            for c := len(grid[0]) - 1; c >=0; c-- {
                if !skip {
                    ans = append(ans, grid[r][c])
                }
                skip = !skip
            }
        }
        forward = !forward
    }

    return ans
}