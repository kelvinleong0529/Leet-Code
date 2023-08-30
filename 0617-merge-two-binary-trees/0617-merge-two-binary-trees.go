func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	var root *TreeNode
	if root1 == nil && root2 == nil {
		return root
	}
	root = new(TreeNode)
	var dfs func(*TreeNode, *TreeNode, *TreeNode)
	dfs = func(root, root1, root2 *TreeNode) {
		var value int
		switch {
		case root1 != nil && root2 != nil:
			value = root1.Val + root2.Val
		case root1 == nil:
			value = root2.Val
		case root2 == nil:
			value = root1.Val
		}
		root.Val = value

		var root1Left, root2Left *TreeNode
		if root1 != nil && root1.Left != nil {
			root1Left = root1.Left
		}
		if root2 != nil && root2.Left != nil {
			root2Left = root2.Left
		}
		if root1Left != nil || root2Left != nil {
			root.Left = new(TreeNode)
			dfs(root.Left, root1Left, root2Left)
		}

		var root1Right, root2Right *TreeNode
		if root1 != nil && root1.Right != nil {
			root1Right = root1.Right
		}
		if root2 != nil && root2.Right != nil {
			root2Right = root2.Right
		}
		if root1Right != nil || root2Right != nil {
			root.Right = new(TreeNode)
			dfs(root.Right, root1Right, root2Right)
		}
	}
	dfs(root, root1, root2)
	return root
}
