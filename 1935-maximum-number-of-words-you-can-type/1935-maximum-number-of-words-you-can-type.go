func canBeTypedWords(text string, brokenLetters string) int {
    var ans int
    var valid = true
    hashMap := make(map[uint8]bool)
    for i := range brokenLetters {
        hashMap[brokenLetters[i]] = true
    }
    for i := range text {
        if string(text[i]) == " " {
            if valid {
                ans++
            }
            valid = true
            continue
        }
        if !valid {
            continue
        }
        _, ok := hashMap[text[i]]
        if ok {
            valid = false
        }
    }
    if valid {
        ans++
    }
    return ans
}