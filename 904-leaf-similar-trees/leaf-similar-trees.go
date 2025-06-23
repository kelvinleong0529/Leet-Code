/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	result1 := make([]int, 0)
	result2 := make([]int, 0)

	var dfs func(*TreeNode, *[]int)
	dfs = func(node *TreeNode, array *[]int) {

		if node == nil {
			return
		}

		dfs(node.Left, array)

		if node.Left == nil && node.Right == nil {
			*array = append(*array, node.Val)
			return
		}

		dfs(node.Right, array)
	}

	dfs(root1, &result1)
	dfs(root2, &result2)

	if len(result1) != len(result2) {
		return false
	}
	for i := 0; i < len(result1); i++ {
		if result1[i] != result2[i] {
			return false
		}
	}
	return true
}