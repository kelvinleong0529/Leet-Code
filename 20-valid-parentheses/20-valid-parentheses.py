class Solution(object):
    def isValid(self, s):
        """
        :type s: str
        :rtype: bool
        """
        da = []
        closed = [")","}","]"]
        hash_map = {")":"(","}":"{","]":"["}
        for i in s:
            if i in closed:
                if len(da) == 0:
                    return False
                if da.pop() != hash_map[i]:
                    return False
            else:
                da.append(i)
        if len(da):
            return False
        else:
            return True
            