class Solution:
    def findRestaurant(self, list1: List[str], list2: List[str]) -> List[str]:
        hash_map = {}
        
        for i in range(len(list1)+len(list2)):
            hash_map[i] = []
        
        least_index_sum = len(list1) + len(list2)
        for i1,v in enumerate(list1):
            if i1 > least_index_sum:
                break
            last_index = least_index_sum - i1 + 1
            if last_index > len(list2):
                last_index = len(list2)
            if v in list2[:last_index]:
                i2 = list2.index(v)
                least_index_sum = i1 + i2
                hash_map[least_index_sum].append(v)
        
        print(hash_map)
    
        for key in hash_map:
            if len(hash_map[key]) > 0:
                return hash_map[key]
        
        return []