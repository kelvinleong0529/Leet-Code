class Solution:
    def trap(self, height: List[int]) -> int:
        if len(height) == 0:
            return 0
        left_index, right_index = 0, len(height) - 1
        left_max, right_max = height[left_index], height[right_index]
        ans = 0
        while left_index < right_index:
            if left_max <= right_max:
                temp = left_max - height[left_index]
                if temp > 0:
                    ans += temp
                left_index += 1
                left_max = max(left_max, height[left_index])
                continue
            else:
                temp = right_max - height[right_index]
                if temp > 0:
                    ans += temp
                right_index -= 1
                right_max = max(right_max, height[right_index])
                continue
        return ans