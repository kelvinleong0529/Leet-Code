class Solution:
    def titleToNumber(self, columnTitle: str) -> int:
        ans = 0
        power = 0
        for i in range(len(columnTitle)-1,-1,-1):
            val = ord(columnTitle[i]) - ord('A') + 1
            ans += val * 26**power
            power += 1
        return ans