func reverseWords(s string) string {
    var ans, temp string

    for _, char := range s {
        if char == ' ' {
            ans += reverseString(temp)
            ans += " "
            temp = ""
        } else {
            temp += string(char)
        }
    }

    ans += reverseString(temp)

    return ans
}

func reverseString(s string) string {
    runes := []rune(s)
    length := len(runes)

    for i := 0; i < length/2; i++ {
        runes[i], runes[length-1-i] = runes[length-1-i], runes[i]
    }

    return string(runes)
}