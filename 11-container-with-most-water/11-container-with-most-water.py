class Solution:
    def maxArea(self, height: List[int]) -> int:
        left_pointer, right_pointer = 0, len(height) - 1
        ans = 0
        while left_pointer < right_pointer:
            area = (right_pointer-left_pointer) * min(height[left_pointer],height[right_pointer])
            ans = max(ans,area)
            if height[left_pointer] <= height[right_pointer]:
                left_pointer += 1
            else:
                right_pointer -= 1
        return ans