func coinChange(coins []int, amount int) int {

    if amount == 0 {
        return 0
    }

    queue := make([]int, 0)
    queue = append(queue, amount)
    visited := make(map[int]bool)
    visited[amount] = true
    steps := 0

    for len(queue) > 0 {
        steps++
        size := len(queue)

        for i := 0; i < size; i++ {

            current := queue[0]
            queue = queue[1:]

            for _, coin := range coins {
                res := current - coin
                if res == 0 {
                    return steps
                }
                if res > 0 && !visited[res] {
                    visited[res] = true
                    queue = append(queue, res)
                }
            }
        }
    }

    return -1
}