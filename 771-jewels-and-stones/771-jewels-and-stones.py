class Solution:
    def numJewelsInStones(self, jewels: str, stones: str) -> int:
        jewels = set(jewels)
        count = 0
        for i in stones:
            if i in jewels:
                count += 1
        return count