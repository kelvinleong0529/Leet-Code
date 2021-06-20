class Solution:
    def uncommonFromSentences(self, s1: str, s2: str) -> List[str]:
        
        def check(target,t1,t2):
            temp = list((t1 + ' ' + t2).split(' '))
            if temp.count(target) == 1:
                return True
            else:
                return False
            
        temp = list(set(list((s1+' '+s2).split(' '))))
        ans = []
        for i in temp:
            if check(i,s1,s2):
                ans.append(i)
        return ans