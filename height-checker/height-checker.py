class Solution:
    def heightChecker(self, heights: List[int]) -> int:
        copy = heights.copy()
        copy.sort()
        for i in range(len(heights)):
            copy[i] = heights[i] - copy[i]
        return len(copy)-copy.count(0)