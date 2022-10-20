class Solution:
    def kthDistinct(self, arr: List[str], k: int) -> str:
        hash_map = {}
        for i in arr:
            if i in hash_map:
                hash_map[i] = False
            else:
                hash_map[i] = True
        
        count = 1
        print(hash_map)
        for i in hash_map:
            if hash_map[i] and count == k:
                return i
            if hash_map[i]:
                count += 1
                
        return ""