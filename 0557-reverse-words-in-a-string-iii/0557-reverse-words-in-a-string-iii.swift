class Solution {
    func reverseWords(_ s: String) -> String {
        var ans = ""
        var temp = ""

        for i in s {
            if i == " " {
                ans += String(temp.reversed())
                ans += " "
                temp = ""
            } else {
                temp += String(i)
            }
        }

        return ans + String(temp.reversed())
    }
}