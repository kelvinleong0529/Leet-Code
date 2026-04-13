func numSquares(n int) int {
    dp := make([]int, n+1)
    
    for i := 1; i <= n; i++ {
        dp[i] = n
    }

    for i := 1; i <= n; i++ {
        for target := 1; target <= i+1; target++ {
            square := target * target
            if square > i {
                break
            }
            dp[i] = min(dp[i], 1 + dp[i - square])
        }
    }

    return dp[n]
}