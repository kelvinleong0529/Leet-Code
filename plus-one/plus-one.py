class Solution:
    def plusOne(self, digits: List[int]) -> List[int]:
        digits = [str(x) for x in digits]
        digits = int("".join(digits))
        digits += 1
        return [int(x) for x in str(digits)]