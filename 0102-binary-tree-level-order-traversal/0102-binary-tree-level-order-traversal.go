/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }
    
    result := make([][]int, 0)
    queue := make([][]*TreeNode,0)
    
    queue = [][]*TreeNode{{root}}
    for {
        if len(queue)  == 0{
            break
        }
        nodes := queue[0]
        queue = queue[1:]
        
        currentLevel := make([]int, 0)
        nextLevel := make([]*TreeNode,0)
        
        for _, node := range nodes {
            currentLevel = append(currentLevel, node.Val)
            if node.Left != nil {
                nextLevel = append(nextLevel, node.Left)
            }
            if node.Right != nil {
                nextLevel = append(nextLevel, node.Right)
            }
        }
        
        result = append(result, currentLevel)
        if len(nextLevel) > 0 {
            queue = append(queue, nextLevel)
        }
    }
    
    return result
}