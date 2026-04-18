func minInsertions(s string) int {
    reversed := ""

    for i := len(s) - 1; i>= 0; i-- {
        reversed += string(s[i])
    }

    dp := make([][]int, len(s) + 1)
    for i := range len(s) + 1 {
        dp[i] = make([]int, len(s) + 1)
    }

    for i := 1; i <= len(s); i++ {
        for j := 1; j <= len(s); j++ {
            if s[i-1] == reversed[j-1] {
                dp[i][j] = dp[i-1][j-1] + 1
            } else {
                dp[i][j] = max(dp[i][j-1], dp[i-1][j])
            }
        }
    }

    return len(s) - dp[len(s)][len(s)]
}