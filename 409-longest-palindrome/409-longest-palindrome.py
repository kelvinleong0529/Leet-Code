class Solution:
    def longestPalindrome(self, s: str) -> int:
        da = {}
        for i in s:
            if i in da:
                da[i] += 1
            else:
                da[i] = 1
        
        ans = 0
        odd = False
        
        for _, value in da.items():
            if value % 2:
                ans += value - 1
                odd = True
            else:
                ans += value
            
        if odd:
            return ans + 1
        else:
            return ans