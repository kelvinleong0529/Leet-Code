class Solution:
    def containsNearbyDuplicate(self, nums: List[int], k: int) -> bool:
        hash_map = {}
        for i,v in enumerate(nums):
            if v in hash_map and i - hash_map[v] <= k:
                    return True
            hash_map[v] = i
        return False
