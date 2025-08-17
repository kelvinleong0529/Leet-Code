/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	ans := -1_001

	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		current := max(node.Val+left+right, node.Val+left, node.Val+right, node.Val)
		if current > ans {
			ans = current
		}
		return node.Val + max(left, right, 0)
	}

	dfs(root)
	return ans
}