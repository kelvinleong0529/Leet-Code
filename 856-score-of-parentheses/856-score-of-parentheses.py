class Solution:
    def scoreOfParentheses(self, s: str) -> int:
        count, ans = 0, 0
        for i,v in enumerate(s):
            if v == "(":
                count += 1
            else:
                count -= 1
                if s[i-1] == "(":
                    ans += 2**(count)
        return ans