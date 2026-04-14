/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isCousins(root *TreeNode, x int, y int) bool {
    queue := make([]*TreeNode, 0)
    queue = append(queue, root)

    for len(queue) > 0 {
        queueLength := len(queue)

        foundX := false
        foundY := false

        for i := 0; i < queueLength; i++ {
            node := queue[0]
            queue = queue[1:]

            if node.Val == x { foundX = true }
            if node.Val == y { foundY = true }

            if node.Left != nil && node.Right != nil {
                if (node.Left.Val == x && node.Right.Val == y) ||
                    (node.Left.Val == y && node.Right.Val == x) {
                        return false
                    } 
            }

            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }

        if foundX && foundY {
            return true
        } else if foundX || foundY {
            return false
        }
    }

    return false
}