class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        hash_map = {}
        for i in range(len(nums)):
            difference = target - nums[i]
            if difference in hash_map:
                hash_map[difference].append(i)
                return hash_map[difference]
            hash_map[nums[i]] = [i]