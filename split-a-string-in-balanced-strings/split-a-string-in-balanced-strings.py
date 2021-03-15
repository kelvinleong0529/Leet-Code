class Solution:
    def balancedStringSplit(self, s: str) -> int:
        count = 0
        sub_count = 0
        for i in s:
            if i == "R":
                count += 1
            else:
                count -= 1
            if not count:
                sub_count += 1
        return sub_count