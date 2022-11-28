func countPoints(rings string) int {
    
    type Color struct {
        blue bool
        green bool
        red bool
    }
    hashMap := make(map[byte]Color)
    
    for i:= 0; i < len(rings); i += 2 {
        color := rings[i]
        num := rings[i+1]
        _, ok := hashMap[num]
        if !ok {
            hashMap[num] = Color{}
        }
        if value, ok := hashMap[num]; ok {
            switch color {
            case 'B':
                value.blue = true
            case 'R':
                value.red = true
            case 'G':
                value.green = true
            }
            hashMap[num] = value
        }
    }
    
    answer := 0
    
    for _, v := range hashMap {
        if v.blue && v.green && v.red {
            answer++
        }
    }
    
    return answer
}
