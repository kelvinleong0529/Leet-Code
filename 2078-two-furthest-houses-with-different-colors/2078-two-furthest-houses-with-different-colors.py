class Solution:
    def maxDistance(self, colors: List[int]) -> int:
        pointer = len(colors) - 1
        ans = 0
        for i in range(len(colors)-1):
            if colors[pointer] != colors[i]:
                if i > len(colors) - 1 - i:
                    if i > ans:
                        ans = i
                else:
                    if len(colors) - i - 1 > ans:
                        ans = len(colors) - i - 1
        return ans