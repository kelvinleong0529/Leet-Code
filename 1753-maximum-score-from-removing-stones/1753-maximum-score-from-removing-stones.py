class Solution:
    def maximumScore(self, a: int, b: int, c: int) -> int:
        heap = [-a,-b,-c]
        heapq.heapify(heap)
        score = 0
        
        while len(heap) > 1:
            first = heapq.heappop(heap) + 1
            second = heapq.heappop(heap) + 1
            if first < 0:
                heapq.heappush(heap,first)
            if second < 0:
                heapq.heappush(heap,second)
            score += 1
        
        return score