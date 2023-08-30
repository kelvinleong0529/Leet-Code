/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	var dfs func(*TreeNode) []int
	dfs = func(node *TreeNode) []int {
		if node == nil {
			return []int{}
		}
		leftNums := dfs(node.Left)
		rightNums := dfs(node.Right)
		nums := append(leftNums, rightNums...)
		return append(nums, node.Val)
	}

	nums1 := dfs(root1)
	nums2 := dfs(root2)

	answer := append(nums1, nums2...)
	sort.Slice(answer, func(i, j int) bool {
		return answer[i] < answer[j]
	})
	return answer
}