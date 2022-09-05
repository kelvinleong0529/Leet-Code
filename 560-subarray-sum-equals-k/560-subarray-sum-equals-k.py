class Solution:
    def subarraySum(self, nums: List[int], k: int) -> int:
        if len(nums) == 1 and nums[0] != k:
            return 0
        hash_map = {0:1}
        total = ans = 0
        for i in nums:
            total += i
            diff = total - k
            if diff in hash_map:
                ans += hash_map[diff]
            if total in hash_map:
                hash_map[total] += 1
            else:
                hash_map[total] = 1
        return ans