class Solution:
    def maxSubsequence(self, nums: List[int], k: int) -> List[int]:
        temp = nums.copy()
        heapq.heapify(temp)
        k = len(nums) - k
        while k > 0:
            nums.remove(heapq.heappop(temp))
            k -= 1
        
        return nums
        
        