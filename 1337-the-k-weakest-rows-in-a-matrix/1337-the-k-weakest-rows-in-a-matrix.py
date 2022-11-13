class Solution:
    
    def count(self, row: List[int]) -> int:
        count = 0
        for i in row:
            if i == 1:
                count += 1
        return count
    
    def kWeakestRows(self, mat: List[List[int]], k: int) -> List[int]:
        hashMap = {}
        ans = []
        
        for i in range(len(mat[0])+1):
            hashMap[i] = []
        
        for i, v in enumerate(mat):
            freq = self.count(v)
            hashMap[freq].append(i)

        for i in hashMap:
            if k == 0:
                return ans
            for j in hashMap[i]:
                if k == 0:
                    return ans
                ans.append(j)
                k -= 1
        return ans