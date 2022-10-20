func groupThePeople(groupSizes []int) [][]int {
    hashMap := make(map[int][]int)
    for i, v := range groupSizes {
        _, ok := hashMap[v]
        if ok {
            hashMap[v] = append(hashMap[v],i)
        } else {
            hashMap[v] = []int {i}
        }
    }
    
    answer := make([][]int,0)
    
    for k1 := range hashMap {
        temp := make([]int,0)
        for _, v2 := range hashMap[k1] {
            temp = append(temp,v2)
            if len(temp) == k1 {
                answer = append(answer, temp)
                temp = []int {}
            }
        }
    }
    
    return answer
}