func maximumWealth(accounts [][]int) int {
    var answer int
    for _, v := range accounts {
        wealth := wealth(v)
        if wealth > answer {
            answer = wealth
        }
    }
    return answer
}

func wealth(account []int) int {
    var answer int
    for _, v := range account {
        answer += v
    }
    return answer
}