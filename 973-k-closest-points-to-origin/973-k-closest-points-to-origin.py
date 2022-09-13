class Solution:
    def kClosest(self, points: List[List[int]], k: int) -> List[List[int]]:
        da = []
        for x,y in points:
            distance = (x**2) + (y**2)
            da.append([distance,x,y])
        
        ans = []
        heapq.heapify(da)
        while k > 0:
            distance,x,y = heapq.heappop(da)
            ans.append([x,y])
            k -= 1
        
        return ans