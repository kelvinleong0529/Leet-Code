class Solution:
    def finalPrices(self, prices: List[int]) -> List[int]:
        stack = [] # [price, index]
        ans = prices.copy()
        
        for i,p in enumerate(prices):
            while stack and p <= stack[-1][0]:
                stackP, stackI = stack.pop()
                ans[stackI] = stackP - p
            stack.append([p,i])
            
        return ans
        