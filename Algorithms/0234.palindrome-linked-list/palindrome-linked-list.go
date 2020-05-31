package problem0234

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type ListNode = kit.ListNode

func isPalindrome(head *ListNode) bool {
	fast,slow:=head,head
	for {
		if slow==nil||slow.Next==nil {
			return true
		}

		slow=slow.Next
		fast=fast.Next.Next
		if fast==nil||fast.Next==nil{
			slow=reverseLink(slow)
			fast=head
			break
		}
	}
	for slow!=nil{
		if slow.Val!=fast.Val{
			return false
		}
		slow=slow.Next
		fast=fast.Next
	}
	return true
}

func reverseLink(head *ListNode) *ListNode{
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