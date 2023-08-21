/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	var answer int
	var recursive func(*TreeNode, int)
	recursive = func(node *TreeNode, level int) {
		if node == nil {
			if level > answer {
				answer = level
			}
			return
		}
		recursive(node.Left, level+1)
		recursive(node.Right, level+1)
	}
	recursive(root, 0)
	return answer
}
