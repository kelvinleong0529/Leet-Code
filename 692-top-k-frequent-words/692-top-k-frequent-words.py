class Solution:
    def topKFrequent(self, words: List[str], k: int) -> List[str]:
        hash_map = {}
        da = [[] for i in range(len(words))]

        for word in words:
            if word in hash_map:
                index = hash_map[word]
                hash_map[word] += 1
                da[index].remove(word)
                da[index+1].append(word)
            else:
                hash_map[word] = 1
                da[1].append(word)
                
        ans = []
        index = len(da) - 1
        
        for i in range(len(da)-1,-1,-1):
            if k == 0:
                break
            if len(da[i]) > 0:
                da[i].sort()
                if len(da[i]) > k:
                    ans.extend(da[i][:k])
                    break
                ans.extend(da[i])
                k -= len(da[i])
        
        return ans