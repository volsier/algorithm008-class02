/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    if headA == nil || headB == nil {
        return nil
    }
    node1,node2 := headA,headB
    for node1 != node2 {
        if node1 != nil {
            node1 = node1.Next
        }else{
            node1 = headB
        }
        if node2 != nil {
            node2 = node2.Next
        }else{
            node2 = headA

        }
    }
    return node1
}