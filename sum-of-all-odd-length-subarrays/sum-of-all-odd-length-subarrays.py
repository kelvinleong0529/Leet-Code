class Solution:
    def sumOddLengthSubarrays(self, arr: List[int]) -> int:
        return sum([arr[i]*(((len(arr)-i)*(i+1)+1)//2) for i in range(len(arr))])