class Solution:
    def findLucky(self, arr: List[int]) -> int:
        temp = list(set(arr))
        temp.sort(reverse = True)
        for i in temp:
            if arr.count(i) == i:
                return i
        else:
            return -1
