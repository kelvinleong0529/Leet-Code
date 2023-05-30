/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
	paths := make([]string, 0)

	var dfs func(*TreeNode, string)
	dfs = func(node *TreeNode, path string) {
		if node == nil {
			return
		}

		var newPath string
		if path == "" {
			newPath = fmt.Sprint(node.Val)
		} else {
			newPath = path + `->` + fmt.Sprint(node.Val)
		}

		if node.Left == nil && node.Right == nil {
			paths = append(paths, newPath)
			return
		}

		dfs(node.Left, newPath)
		dfs(node.Right, newPath)
	}

	dfs(root, "")
	return paths
}