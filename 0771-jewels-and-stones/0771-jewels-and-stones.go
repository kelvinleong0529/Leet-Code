func numJewelsInStones(jewels string, stones string) int {
    jewelsMap := make(map[rune]bool)
    var answer int
    for _, jewel := range jewels {
        jewelsMap[jewel] = true
    }
    for _, stone := range stones {
        if _, ok := jewelsMap[stone]; ok {
            answer++
        }
    }
    return answer
}