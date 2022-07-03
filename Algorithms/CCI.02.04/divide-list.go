package problemCCI0204

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

func partition(head *ListNode, x int) *ListNode {
    first:=&ListNode{}
    second:=&ListNode{}
    res1:=first
    res2:=second
    for head!=nil {
        if head.Val<x{
            first.Next=&ListNode{
                head.Val,
                nil,
            }
            first=first.Next
        }else{
            second.Next=&ListNode{
                head.Val,
                nil,
            }
            second=second.Next
        }
        head=head.Next
    }
    first.Next=res2.Next
    return res1.Next
}