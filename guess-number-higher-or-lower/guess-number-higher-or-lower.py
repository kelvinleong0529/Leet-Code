# The guess API is already defined for you.
# @param num, your guess
# @return -1 if my number is lower, 1 if my number is higher, otherwise return 0
# def guess(num: int) -> int:

class Solution:
    def guessNumber(self, n: int) -> int:
        start = 1
        while 1:
            if not guess(int((start+n)/2)) :
                return int((start+n)/2)
            elif guess(int((start+n)/2)) == 1:
                start = int((start+n)/2) + 1
            else:
                n = int((start+n)/2) - 1
                