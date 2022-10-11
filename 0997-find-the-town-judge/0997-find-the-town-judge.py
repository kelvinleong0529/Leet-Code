class Solution:
    def findJudge(self, n: int, trust: List[List[int]]) -> int:
        
        if len(trust) == 0 and n == 1:
            return 1
        
        hash_map = {}
        
        for t in trust:
            p1, p2 = t[0], t[-1]
            if p1 in hash_map:
                hash_map[p1][0] += 1
            else:
                hash_map[p1] = [1,0]
            if p2 in hash_map:
                hash_map[p2][-1] += 1
            else:
                hash_map[p2] = [0,1]
            
        for key in hash_map:
            if hash_map[key][0] == 0 and hash_map[key][-1] == n-1:
                return key
        
        return -1