/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
	var dfs func(*TreeNode, bool, int) int
	dfs = func(node *TreeNode, isLeft bool, sum int) int {
		if node == nil {
			return 0
		}

		if node.Left == nil && node.Right == nil && isLeft {
			return node.Val
		}

		return dfs(node.Left, true, sum) + dfs(node.Right, false, sum)
	}

	return dfs(root.Left, true, 0) + dfs(root.Right, false, 0)
}