# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def rangeSumBST(self, root: Optional[TreeNode], low: int, high: int) -> int:

        def dfs(root):
            
            total = 0
            if not root:
                return 0

            total += dfs(root.left)
            if root.val >= low and root.val <= high:
                total += root.val

            total += dfs(root.right)
            return total
        
        return dfs(root)