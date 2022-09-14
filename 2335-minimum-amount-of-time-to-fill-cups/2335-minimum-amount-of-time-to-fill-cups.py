class Solution:
    def fillCups(self, amount: List[int]) -> int:
        if sum(amount) == 0:
            return 0
        amount = [-a for a in amount]
        heapq.heapify(amount)
        time = 0
        while amount:
            if len(amount) == 1:
                time += abs(amount[0])
                break
            first = heapq.heappop(amount) + 1
            second = heapq.heappop(amount) + 1
            if first < 0:
                heapq.heappush(amount,first)
            if second < 0:
                heapq.heappush(amount,second)
            time += 1
            
        return time