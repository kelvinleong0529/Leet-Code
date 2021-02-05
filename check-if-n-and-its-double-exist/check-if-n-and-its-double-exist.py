class Solution:
    def checkIfExist(self, arr: List[int]) -> bool:
        if arr.count(0) > 1:
            return 1
        arr.sort(reverse=True)
        for i in arr:
            if not i:
                continue
            if i%2:
                continue
            if i/2 in arr:
                return 1
        return 0