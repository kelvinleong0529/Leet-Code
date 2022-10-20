func checkIfPangram(sentence string) bool {
    hashMap := make(map[byte]int)
    for i:= 'a'; i <= 'z'; i++ {
        hashMap[[]byte(string(i))[0]] = 0
    }
    for i := range sentence {
        hashMap[sentence[i]]++
    }
    for _, v := range hashMap {
        if v == 0 {
            return false
        }
    }
    return true
}