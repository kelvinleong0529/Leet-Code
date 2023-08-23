func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
    var answer int
    for _, v := range hours {
        if v >= target {
            answer++
        }
    }
    return answer
}