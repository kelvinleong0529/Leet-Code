class Solution:
    def carPooling(self, trips: List[List[int]], capacity: int) -> bool:
        trips = [[f,t,n] for n,f,t in trips]
        heapq.heapify(trips)
        stop = []
        heapq.heapify(stop)
        location = 0
        
        while len(trips) > 0:
            while stop and stop[0][0] == location:
                _, stackN = heapq.heappop(stop)
                capacity += stackN
            while trips and trips[0][0] == location:
                _, stackS, stackN = heapq.heappop(trips)
                if capacity < stackN:
                    return False
                capacity -= stackN
                heapq.heappush(stop,[stackS,stackN])
            location += 1
        
        return True