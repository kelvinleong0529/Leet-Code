class Solution:
    def reverseWords(self, s: str) -> str:
        ans, temp = "", ""
        for i in s:
            if i == " ":
                ans += temp[::-1]
                ans += " "
                temp = ""
            else:
                temp += i
        
        return ans + temp[::-1]