/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum(head *ListNode) int {
    list := make([]int, 0)
    var length int
    
    for {
        if head == nil {
            break
        }
        list = append(list, head.Val)
        head = head.Next
        length++
    }
    
    var maxSum int
    for i:=0; i< length/2; i++ {
        sum := list[i] + list[length-1-i]
        if sum > maxSum {
            maxSum = sum
        }
    }
    
    return maxSum
}