/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
	var recursive func(*TreeNode, int) bool
	recursive = func(node *TreeNode, sum int) bool {
		if node == nil {
			return false
		}
		if node.Left == nil && node.Right == nil {
			return sum+node.Val == targetSum
		}
		return recursive(node.Left, sum+node.Val) || recursive(node.Right, sum+node.Val)
	}
	return recursive(root, 0)
}