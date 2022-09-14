class Solution:
    def minStoneSum(self, piles: List[int], k: int) -> int:
        piles = [-p for p in piles]
        heapq.heapify(piles)
        
        while k > 0:
            cur = heapq.heappop(piles)
            heapq.heappush(piles,cur//2)
            k -= 1
        
        return -sum(piles)
