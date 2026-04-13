func coinChange(coins []int, amount int) int {
    dp := make([]int, amount + 1)

    for i := 1; i <= amount; i++ {
        dp[i] = amount+1
    }

    for i := 1; i <= amount; i++ {
        for _, c := range coins {
            res := i - c
            if res < 0 {
                continue
            }
            dp[i] = min(dp[i], 1 + dp[res])
        }
    }

    // if dp[amount] == amount + 1, no solution
    if dp[amount] > amount {
        return -1
    }
    return dp[amount]
}