/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func postorder(root *Node) []int {
    
    result := make([]int,0)
    
    var postOrder func(*Node)
    postOrder = func(node *Node) {
        if node == nil {
            return
        }
        
        for i := range node.Children {
            postOrder(node.Children[i])
        }
        
        result = append(result, node.Val)
    }
    
    postOrder(root)
    
    return result
    
}