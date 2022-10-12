class Solution:
    def licenseKeyFormatting(self, s: str, k: int) -> str:
        ans = ""
        temp = ""
        for i in range(len(s)-1,-1,-1):
            if len(temp) == k:
                ans += "-"
                ans += temp
                temp = ""
            if s[i] == "-":
                continue
            temp += s[i].upper()
        
        if len(temp) > 0:
            ans += "-"
            ans += temp
        
        return ans[1:][::-1]