func isAcronym(words []string, s string) bool {
    if len(words) != len(s) {
        return false
    }
    for i, v := range words {
        if v[0] != s[i] {
            return false
        }
    }
    return true
}