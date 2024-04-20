/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isUnivalTree(root *TreeNode) bool {
    var dfs func(*TreeNode, int) bool
    dfs = func(node *TreeNode, val int) bool {
        if node == nil {
            return true;
        }
        
        if node.Val != val {
            return false;
        }
        
        return dfs(node.Left, val) && dfs(node.Right, val)
    }
    
    return dfs(root, root.Val)
}