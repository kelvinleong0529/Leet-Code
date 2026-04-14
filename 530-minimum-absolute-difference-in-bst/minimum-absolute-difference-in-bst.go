/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getMinimumDifference(root *TreeNode) int {
    ans := 100001
    var previous *TreeNode

    var dfs func(*TreeNode)
    dfs = func(node *TreeNode) {
        if node == nil {
            return
        }
        dfs(node.Left)
        if previous != nil {
            ans = min(ans, abs(node.Val - previous.Val))
        }
        previous = node
        dfs(node.Right)
    }
    dfs(root)

    return ans
}

func abs(n int) int {
    if n < 0 {
        return -n
    }
    return n
}