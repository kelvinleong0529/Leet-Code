/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
    result := make([]int, 0)
    var preOrder func(*Node)
    
    preOrder = func(node *Node) {
        if node == nil {
            return
        }
        
        result = append(result, node.Val)
        
        for i := range node.Children {
            preOrder(node.Children[i])
        }
    }
    
    preOrder(root)
    
    return result
}