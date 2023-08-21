/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	var recursiveFunc func(*TreeNode, *TreeNode) bool
	recursiveFunc = func(p, q *TreeNode) bool {
		if p == nil && q == nil {
			return true
		}
		if p == nil || q == nil {
			return false
		}
		return p.Val == q.Val && recursiveFunc(p.Left, q.Left) && recursiveFunc(p.Right, q.Right)
	}
	return recursiveFunc(p, q)
}