func areOccurrencesEqual(s string) bool {
    hashMap := make(map[byte]int)
    
    for i := range s {
        _, ok := hashMap[s[i]]
        if ok {
            hashMap[s[i]]++
        } else {
            hashMap[s[i]] = 1
        }
    }
    
    occurence := 0
    
    for _, v := range hashMap {
        if occurence == 0 {
            occurence = v
            continue
        }
        if v != occurence {
            return false
        }
    }
    
    return true
}