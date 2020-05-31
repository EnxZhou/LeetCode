package problem0142

import "github.com/aQuaYi/LeetCode-in-Go/kit"

// ListNode is pre-defined...
type ListNode = kit.ListNode

// https://leetcode.com/problems/linked-list-cycle-ii/discuss/44793/O(n)-solution-by-using-two-pointers-without-change-anything
func detectCycle1(head *ListNode) *ListNode {
	if head==nil{
		return nil
	}
	nodeMap:=make(map[*ListNode]int)
	new:=head
	for {
		if new==nil{
			return nil
		}
		if _,ok:=nodeMap[new];ok{
			return new
		}else{
			nodeMap[new]=new.Val
			new=new.Next
		}
	}
}

func detectCycle(head *ListNode) *ListNode {
	fast,slow:=head,head
	for {
		if fast==nil||fast.Next==nil{
			return nil
		}
		fast=fast.Next.Next
		slow=slow.Next
		if fast==slow{
			break
		}
	}
	fast=head
	for fast!=slow{
		fast=fast.Next
		slow=slow.Next
	}
	return fast
}