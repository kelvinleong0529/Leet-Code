class Solution:
    def largestSumAfterKNegations(self, A: List[int], K: int) -> int:
        for i in range(K):
            if min(A):
                A[A.index(min(A))] = -min(A)
            else:
                break
        return sum(A)
