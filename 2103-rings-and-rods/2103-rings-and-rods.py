class Solution:
    def countPoints(self, rings: str) -> int:
        ans = 0
        hash_map = {}
        for i in range(10):
            hash_map[i] = [0,0,0]
        color_map = {
            "R": 0,
            "G": 1,
            "B": 2
        }
        for i in range(0,len(rings),2):
            color = rings[i]
            index = int(rings[i+1])
            if hash_map[index][color_map[color]] == 0:
                hash_map[index][color_map[color]] = 1
        
        for key in hash_map:
            if sum(hash_map[key]) == 3:
                ans += 1
        
        return ans