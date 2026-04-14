func lexicalOrder(n int) []int {
    ans := make([]int, 0)

    var dfs func(str string)
    dfs = func(str string) {
        for i := 0; i <= 9; i++ {
            intString := str + strconv.Itoa(i)
            newInt, _ := strconv.Atoi(intString)
            if newInt > n {
                return
            }
            if newInt > 0 {
                ans = append(ans, newInt)
                dfs(intString)
            }
        }
    }

    dfs("")

    return ans
}