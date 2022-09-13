class Solution:
    def maxPower(self, s: str) -> int:
        ans = 0
        prev = None
        count = 1
        for i in s:
            if i == prev:
                count += 1
            else:
                if count > ans:
                    ans = count
                count = 1
                prev = i
        if count > ans:
            return count
        else:
            return ans