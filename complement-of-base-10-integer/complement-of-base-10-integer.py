import math
class Solution:
    def bitwiseComplement(self, N: int) -> int:
        power = 1
        while (pow(2,power)-1 < N):
            power +=1
        return pow(2,power)-N-1
