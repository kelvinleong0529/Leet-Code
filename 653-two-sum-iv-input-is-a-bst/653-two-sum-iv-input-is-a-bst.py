# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def findTarget(self, root: Optional[TreeNode], k: int) -> bool:
        
        remaining = set()
        
        def dfs(root):
            if not root:
                return False
            if root.val in remaining:
                return True
            
            diff = k - root.val
            remaining.add(diff)
            
            res = dfs(root.left) or dfs(root.right)
            return res
        
        return dfs(root)