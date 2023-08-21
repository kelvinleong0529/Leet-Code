/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var recursive func(*TreeNode, *TreeNode) bool
	recursive = func(p, q *TreeNode) bool {
		if p != nil && q != nil {
			return p.Val == q.Val && recursive(p.Left, q.Right) && recursive(p.Right, q.Left)
		}
		return p == q
	}
	return recursive(root.Left, root.Right)
}
