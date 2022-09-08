class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        ans = ""
        ans_len = 0
        for i in s:
            while i in ans:
                ans = ans[1:]
            ans += i
            if len(ans) > ans_len:
                ans_len = len(ans)
        
        return ans_len
        