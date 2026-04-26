/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func tree2str(root *TreeNode) string {
	ans := make([]string, 0)

	var dfs func(*TreeNode)

	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		ans = append(ans, "(")
		ans = append(ans, strconv.Itoa(node.Val))

		if node.Left == nil && node.Right != nil {
			ans = append(ans, "()")
		}

		dfs(node.Left)
		dfs(node.Right)

		ans = append(ans, ")")
	}

	dfs(root)

	return strings.Join(ans[1:len(ans)-1], "")
}