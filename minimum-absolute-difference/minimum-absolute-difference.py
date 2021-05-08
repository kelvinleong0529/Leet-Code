class Solution:
    def minimumAbsDifference(self, arr: List[int]) -> List[List[int]]:
        arr.sort()
        mindiff = min( [arr[i+1]-arr[i] for i in range(len(arr)-1)])
        ans = [[arr[i],arr[i+1]] for i in range(len(arr)-1) if arr[i+1]-arr[i]==mindiff]
        return ans