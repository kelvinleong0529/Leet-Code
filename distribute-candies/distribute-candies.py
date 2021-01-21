class Solution(object):
    def distributeCandies(self, candies):
        """
        :type candies: List[int]
        :rtype: int
        """
        candyType = set(candies)
        return min(len(candyType), len(candies)//2)
