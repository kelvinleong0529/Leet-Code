class Solution:
    def uniqueOccurrences(self, arr: List[int]) -> bool:
        hash_map = {}
        da = [[] for i in range(len(arr)+1)]
        
        for i in arr:
            if i in hash_map:
                index = hash_map[i]
                da[index].remove(i)
                da[index+1].append(i)
                hash_map[i] += 1
            else:
                hash_map[i] = 1
                da[1].append(i)
        
        for i in da:
            if len(i) > 1:
                return False
        
        return True