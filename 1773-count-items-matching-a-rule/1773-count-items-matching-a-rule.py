class Solution:
    def countMatches(self, items: List[List[str]], ruleKey: str, ruleValue: str) -> int:
        hash_map = {
            "type":0,
            "color":1,
            "name":2
        }
        
        index = hash_map[ruleKey]
        ans = 0
        
        for item in items:
            if item[index] == ruleValue:
                ans += 1
        
        return ans