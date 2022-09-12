# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def maxDepth(self, root: Optional[TreeNode]) -> int:
        
        if not root:
            return 0
        
        def dfs(root):
            if not root:
                return None
            left = dfs(root.left)
            right = dfs(root.right)
            if left or right:
                if not left:
                    return 1 + right
                if not right:
                    return 1 + left
                if left >= right:
                    return 1 + left
                else:
                    return 1 + right
            return 1
        
        return dfs(root)