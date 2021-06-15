# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def hasPathSum(self, root: TreeNode, targetSum: int) -> bool:
        if not root:
            return False 
        return self.dfs(root, targetSum, 0)
    
    def dfs(self, root, targetSum, rSum): 
        # base case 
        if not root:
            return False
        
        # Logic 
        rSum += root.val 
        if self.dfs(root.left, targetSum, rSum):
            return True 
        if self.dfs(root.right, targetSum, rSum):
            return True 
    
        if not root.left and not root.right and targetSum == rSum:
            return True 
        
        return False