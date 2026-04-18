/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {

    dummy := &ListNode{
        Next: head,
    }
    
    left := dummy
    right := head
    for range n {
        right = right.Next
    }

    for right != nil {
        left = left.Next
        right = right.Next
    }

    left.Next = left.Next.Next

    return dummy.Next
}