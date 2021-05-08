class Solution:
    def relativeSortArray(self, arr1: List[int], arr2: List[int]) -> List[int]:
        ans = []
        for i in arr2:
            for j in range(arr1.count(i)):
                ans.append(i)
        temp = [i for i in arr1+arr2 if i not in arr1 or i not in arr2]
        temp.sort()
        return ans + temp