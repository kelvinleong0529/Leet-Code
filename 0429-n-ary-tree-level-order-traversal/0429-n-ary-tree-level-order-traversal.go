/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
    if root == nil {
        return [][]int{}
    }
    
    result := make([][]int, 0)
    queue := make([][]*Node,0)
    
    queue = [][]*Node{{root}}
    
    for {
        if len(queue) == 0 {
            break
        }
        
        nodes := queue[0]
        queue = queue[1:]

        currentLevel := make([]int, 0)
        nextLevelNodes := make([]*Node, 0)
        for _, node := range nodes {
            currentLevel = append(currentLevel, node.Val)
            for _, child := range node.Children {
                nextLevelNodes = append(nextLevelNodes, child)
            }
        }
        
        result = append(result, currentLevel)
        if len(nextLevelNodes) > 0 {
            queue = append(queue, nextLevelNodes)
        }
    }
    
    return result
        
}