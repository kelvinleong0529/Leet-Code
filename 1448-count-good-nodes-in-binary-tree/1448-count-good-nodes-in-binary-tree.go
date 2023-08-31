/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func goodNodes(root *TreeNode) int {
	var answer int
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, max int) {
		if node == nil {
			return
		}
		if node.Val >= max {
			answer++
			max = node.Val
		}
		dfs(node.Left, max)
		dfs(node.Right, max)
	}
	dfs(root, root.Val)
	return answer
}