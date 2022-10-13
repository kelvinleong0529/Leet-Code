class Solution:
    def checkZeroOnes(self, s: str) -> bool:
        max_ones, max_zeros = 0, 0
        ones, zeros = 0, 0
        previous = None
        
        for i in s:
            if previous is None:
                previous = i
            if i != previous:
                ones, zeros = 0, 0
                previous = i
            if i == "1":
                ones += 1
                if ones > max_ones:
                    max_ones = ones
            else:
                zeros += 1
                if zeros > max_zeros:
                    max_zeros = zeros
            
        return True if max_ones > max_zeros else False
                