func repeatedCharacter(s string) byte {
    hashMap := make(map[byte]bool)
    
    for i := range s {
        _, ok := hashMap[s[i]]
        if ok {
            return s[i]
        } else {
            hashMap[s[i]] = true
        }
    }
    return 0
}