class Solution:
    def nextGreaterElement(self, nums1: List[int], nums2: List[int]) -> List[int]:
        hash_map = {}
        for i in range(len(nums2)-1,-1,-1):
            num = nums2[i]
            if i == len(nums2) - 1:
                hash_map[num] = -1
                continue
            next_num = nums2[i+1]
            if num < next_num:
                hash_map[num] = next_num
            else:
                while num > hash_map[next_num]:
                    if hash_map[next_num] == -1:
                        break
                    next_num = hash_map[next_num]
                hash_map[num] = hash_map[next_num]
            
        ans = []
        for n in nums1:
            ans.append(hash_map[n])
        
        return ans