func restoreString(s string, indices []int) string {
    shuffled := make([]byte, len(s))

    for i := 0; i < len(s); i++ {
        shuffled[indices[i]] = s[i]
    }

    return string(shuffled)
}