class Solution:
    def twoSum(self, numbers: List[int], target: int) -> List[int]:
        point = 0
        ans = []
        while point<len(numbers):
            temp = target - numbers[point]
            if temp in numbers[point+1:]:
                ans.append(point+1)
                ans.append(point+numbers[point+1:].index(temp)+2)
                break
            point += 1
        return ans
                