func countSeniors(details []string) int {
    var answer int
    for _, v := range details {
        ageString := v[11:13]
        age := 0
        fmt.Sscanf(ageString,"%d",&age)
        if age > 60 {
            answer++
        }
    }
    return answer
}