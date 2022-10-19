func countPoints(rings string) int {
    hashMap := make(map[byte][]bool)
    
    colorMap := map[byte]int {'B':0,'G':1,'R':2}
    
    for i:= 0; i < len(rings); i += 2 {
        color := rings[i]
        num := rings[i+1]
        index := colorMap[color]
        _, ok := hashMap[num]
        if !ok {
            hashMap[num] = []bool {false, false, false}
        }
        hashMap[num][index] = true
    }
    
    answer := 0
    
    for _, v := range hashMap {
        if v[0] && v[1] && v[2] {
            answer++
        }
    }
    
    return answer
}
