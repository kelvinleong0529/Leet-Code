class Solution:
    def minimumOperations(self, nums: List[int]) -> int:
        temp_heap = []
        for n in nums:
            if n > 0:
                temp_heap.append(n)
        nums = temp_heap
        ans = 0
        heapq.heapify(nums)
        
        while nums:
            temp_heap = []
            smallest = heapq.heappop(nums)
            for n in nums:
                diff = n - smallest
                if diff > 0:
                    temp_heap.append(diff)
                nums = temp_heap
                heapq.heapify(nums)
            ans += 1
            
        return ans