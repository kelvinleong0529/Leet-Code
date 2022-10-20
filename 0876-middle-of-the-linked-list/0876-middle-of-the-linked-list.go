/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    firstPointer := head
    secondPointer := head
    for {
        if firstPointer == nil || firstPointer.Next == nil {
            return secondPointer
        }
        firstPointer = firstPointer.Next
        if firstPointer != nil {
            firstPointer = firstPointer.Next
        }
        secondPointer = secondPointer.Next
    }
    return nil
}