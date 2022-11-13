func arrayStringsAreEqual(word1 []string, word2 []string) bool {
    var s1,s2 string
    for i := range word1 {
        s1 = s1 + word1[i]
    }
    for i := range word2 {
        s2 = s2 + word2[i]
    }
    return s1 == s2
}