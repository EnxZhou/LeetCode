package problem0206

func reverseList(head *ListNode) *ListNode {
	if head==nil||head.Next==nil{
		return head
	}
	var pre *ListNode
	for head!=nil{
		tmp:=head.Next
		head.Next=pre
		pre=head
		head=tmp
	}
	return pre
}

// ListNode 是链接节点
type ListNode struct {
	Val  int
	Next *ListNode
}
