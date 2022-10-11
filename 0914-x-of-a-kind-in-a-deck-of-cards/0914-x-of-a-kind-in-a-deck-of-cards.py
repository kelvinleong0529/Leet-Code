class Solution:
    def hasGroupsSizeX(self, deck: List[int]) -> bool:
        
        def HCF(x:int,y:int) -> int:
            if x == y:
                return x
            if x > y:
                big = x
                small = y
            else:
                big = y
                small = x
            for i in range(small,0,-1):
                if not small%i and not big%i:
                    return i
        
        hash_map = {}
        for card in deck:
            if card in hash_map:
                hash_map[card] += 1
            else:
                hash_map[card] = 1
        
        previous = None
        
        for key in hash_map:
            if previous == None:
                previous = hash_map[key]
                if previous == 1:
                    return False
                continue
            hcf = HCF(previous,hash_map[key])
            if hcf == 1:
                return False
            previous = hcf
        
        return True