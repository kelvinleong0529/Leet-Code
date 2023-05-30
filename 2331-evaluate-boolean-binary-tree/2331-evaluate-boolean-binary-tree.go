/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func evaluateTree(root *TreeNode) bool {
    var dfs func(*TreeNode) bool 
	dfs = func(node *TreeNode) bool {
		if node.Left == nil && node.Right == nil {
			if node.Val == 0 {
				return false
			} else {
				return true
			}
		}

		left := dfs(node.Left)
		right := dfs(node.Right)

		if node.Val == 2 {
			return left || right
		} else {
			return left && right
		}
	}

	return dfs(root)
}