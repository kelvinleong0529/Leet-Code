class Solution:
    def findRelativeRanks(self, score: List[int]) -> List[str]:
        arr = [(-s,i) for i,s in enumerate(score)]
        ans = [-1] * len(score)
        heapq.heapify(arr)
        count = 1
        while count <= 3 and arr:
            _, stackI = heapq.heappop(arr)
            if count == 1:
                ans[stackI] = "Gold Medal"
            elif count == 2:
                ans[stackI] = "Silver Medal"
            else:
                ans[stackI] = "Bronze Medal"
            count += 1
        
        while arr:
            _, stackI = heapq.heappop(arr)
            ans[stackI] = str(count)
            count += 1
        
        return ans