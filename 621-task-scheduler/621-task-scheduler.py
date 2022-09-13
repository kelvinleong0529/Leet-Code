class Solution:
    def leastInterval(self, tasks: List[str], n: int) -> int:
        hash_table = {}
        
        for t in tasks:
            if t in hash_table:
                hash_table[t] += 1
            else:
                hash_table[t] = 1
        
        max_heap = [-i for i in hash_table.values()]
        heapq.heapify(max_heap)
        time = 0
        q = deque()
        
        while q or max_heap:
            time += 1
            
            if max_heap:
                count = 1 + heapq.heappop(max_heap)
                if count:
                    q.append([count,time+n])
            if q and q[0][1] == time:
                heapq.heappush(max_heap,q.popleft()[0])
            
        return time