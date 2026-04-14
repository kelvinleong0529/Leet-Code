func longestPalindrome(s string) string {
    longest := 0
    var ans string

    for i := 0; i < len(s); i++ {

        l, r := i, i

        // odd case
        for l >= 0 && r < len(s) && s[l] == s[r] {
            length := r - l + 1
            if length > longest {
                ans = s[l:r+1]
                longest = length
            }
            l--
            r++
        }

        l, r = i, i + 1
        // even case
        for l >= 0 && r < len(s) && s[l] == s[r] {
            length := r - l + 1
            if length > longest {
                ans = s[l:r+1]
                longest = length
            }
            l--
            r++
        }
    }

    return ans
}