class Solution:
    def twoSum(self, numbers: List[int], target: int) -> List[int]:
        left_pointer = 0
        right_pointer = len(numbers) - 1
        while True:
            temp = numbers[left_pointer] + numbers[right_pointer]
            if temp == target:
                return [left_pointer + 1,right_pointer + 1]
            if temp > target:
                right_pointer -= 1
            else:
                left_pointer += 1