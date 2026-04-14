/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    ans := make([][]int, 0)

    if root == nil {
        return ans
    }
    
    queue := make([]*TreeNode, 0)
    queue = append(queue, root)

    // repeat as long as queue is not empty
    for len(queue) > 0 {
        queueLength := len(queue)
        row := make([]int, 0)

        for i := 0; i < queueLength; i++ {
            node := queue[0]
            queue = queue[1:]

            row = append(row, node.Val)
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }

        if len(row) > 0 {
            ans = append(ans, row)
        }
    }

    return ans
}