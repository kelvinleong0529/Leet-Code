/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 **/
func sumEvenGrandparent(root *TreeNode) int {
	var answer int
	var dfs func(*TreeNode, bool, bool)
	dfs = func(node *TreeNode, grandparent, parent bool) {
		if node == nil {
			return
		}
		if grandparent {
			answer += node.Val
		}
		isEven := node.Val%2 == 0
		dfs(node.Left, parent, isEven)
		dfs(node.Right, parent, isEven)
	}
	dfs(root, false, false)
	return answer
}