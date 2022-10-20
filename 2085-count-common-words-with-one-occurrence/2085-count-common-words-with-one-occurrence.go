func countWords(words1 []string, words2 []string) int {
    hashMap := make(map[string][]int)
    
    update := func(word []string, index int) {
        for i := range word {
            _, ok := hashMap[word[i]]
            if !ok {
                hashMap[word[i]] = []int{0,0}
            }
            hashMap[word[i]][index]++
        }
    }
    
    update(words1, 0)
    update(words2, 1)
    
    answer := 0
    for _, v := range hashMap {
        if v[0] == 1 && v[1] == 1 {
            answer++
        }
    }
    return answer
}