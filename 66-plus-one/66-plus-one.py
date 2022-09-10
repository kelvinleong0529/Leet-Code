class Solution:
    def plusOne(self, digits: List[int]) -> List[int]:
        digits = digits[::-1]
        carry = 0
        
        for i in range(len(digits)):
            if i == 0:                
                result = digits[i] + 1
                carry = result // 10
                digits[i] = result % 10
            else:
                result = carry + digits[i]
                carry = result // 10
                digits[i] = result % 10
        
        if carry:
            digits.append(1)
        
        return digits[::-1]