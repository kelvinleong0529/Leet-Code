class Solution:
    def longestValidParentheses(self, s: str) -> int:

        ans1, ans2 = 0, 0
        l, r = 0, 0

        for i in s:
            if i == "(":
                l += 1
            else:
                r += 1
            if l == r:
                if l*2 > ans1:
                    ans1 = l*2
            if r > l:
                r, l = 0, 0
                
        l, r = 0, 0
        s = s[::-1]
        for i in s:
            if i == "(":
                l += 1
            else:
                r += 1
            if l == r:
                if l*2 > ans2:
                    ans2 = l*2
            if l > r:
                r, l = 0, 0
                
        if ans1 > ans2:
            return ans1
        else:
            return ans2