package problemCCI0201

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// ListNode 是链接节点
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeRepeatNode(head *ListNode) *ListNode {
	if head==nil{
        return nil
    }
    current:=head
    for current!=nil{
        p:=current
        for p.Next!=nil {
            if p.Next.Val==current.Val{
                p.Next=p.Next.Next
            }else{
                p=p.Next
            }
        }
        current=current.Next
    }
    return head
}
