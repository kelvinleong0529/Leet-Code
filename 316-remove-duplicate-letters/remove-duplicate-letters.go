func removeDuplicateLetters(s string) string {
    lastIndex := [26]int{}
    for i, c := range s {
        lastIndex[c-'a'] = i
    }

    inStack := [26]bool{}
    stack := []byte{}

    for i := 0; i < len(s); i++ {
        c := s[i]
        if inStack[c-'a'] {
            continue  // already in result, skip
        }
        // pop if current char is smaller AND top char appears later
        for len(stack) > 0 {
            top := stack[len(stack)-1]
            if c < top && lastIndex[top-'a'] > i {
                stack = stack[:len(stack)-1]
                inStack[top-'a'] = false
            } else {
                break
            }
        }
        stack = append(stack, c)
        inStack[c-'a'] = true
    }

    return string(stack)
}