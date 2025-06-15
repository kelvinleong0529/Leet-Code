func isValidBST(root *TreeNode) bool {
	var recursive func(*TreeNode, int, int) bool

	recursive = func(node *TreeNode, min int, max int) bool {
		if node == nil {
			return true
		}

		var current bool
		current = node.Val < max && node.Val > min

		return current && recursive(node.Left, min, node.Val) && recursive(node.Right, node.Val, max)
	}

	return recursive(root.Left, -1<<31-1, root.Val) && recursive(root.Right, root.Val, 1<<31+1)
}