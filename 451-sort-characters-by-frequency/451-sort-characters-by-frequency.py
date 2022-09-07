class Solution(object):
    def frequencySort(self, s):
        """
        :type s: str
        :rtype: str
        """
        hash_map = {}
        hash_array = [[] for i in range(len(s)+1)]
        for i in s:
            if i in hash_map:
                index = hash_map[i]
                hash_array[index].remove(i)
                hash_array[index+1].append(i)
                hash_map[i] += 1
            else:
                hash_map[i] = 1
                hash_array[1].append(i)
        print(hash_map)
        print(hash_array)
        ans = ""
        for i in range(len(hash_array)-1):
            for j in hash_array[len(hash_array)-1-i]:
                for k in range(len(hash_array)-1-i):
                    ans += j
        return ans