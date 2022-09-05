class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        buy, profit = prices[0], 0
        for i in prices:
            earn = i - buy
            if earn > profit:
                profit = earn
            if i < buy:
                buy = i
        return profit