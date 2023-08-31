/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	var dfs func([]int) *TreeNode
	dfs = func(nums []int) *TreeNode {
		mid := len(nums) / 2
		root := &TreeNode{
			Val: nums[mid],
		}
		if len(nums[:mid]) > 0 {
			root.Left = dfs(nums[:mid])
		}
		if len(nums[mid+1:]) > 0 {
			root.Right = dfs(nums[mid+1:])
		}
		return root
	}
	return dfs(nums)
}