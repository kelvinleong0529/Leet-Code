class Solution:
    def countOdds(self, low: int, high: int) -> int:
        if high-low == 1:
            return 1
        elif not (low%2 or high%2):
            return int((high-low)/2)
        else:
            return int((high-low)/2)+1
