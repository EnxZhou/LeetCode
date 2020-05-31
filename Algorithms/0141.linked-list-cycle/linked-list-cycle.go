package problem0141

import "github.com/aQuaYi/LeetCode-in-Go/kit"

// ListNode is pre-defined...
type ListNode = kit.ListNode

func hasCycle(head *ListNode) bool {
	if head==nil||head.Next==nil{
		return false
	}
	var fast,slow *ListNode
	fast=head.Next
	slow=head
	for fast!=slow {
		if fast==nil||fast.Next==nil{
			return false
		}
		fast=fast.Next.Next
		slow=slow.Next
	}
	return true
}
