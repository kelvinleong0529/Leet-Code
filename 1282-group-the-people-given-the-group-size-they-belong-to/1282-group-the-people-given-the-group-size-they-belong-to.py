class Solution:
    def groupThePeople(self, groupSizes: List[int]) -> List[List[int]]:
        hash_map = {}
        ans = []
        
        def update(k:int) -> None:
            if len(hash_map[k]) == k:
                ans.append(hash_map[k])
                hash_map[k] = []
        
        for i,v in enumerate(groupSizes):
            if v in hash_map:
                hash_map[v].append(i)
            else:
                hash_map[v] = [i]
            update(v)
        
        return ans