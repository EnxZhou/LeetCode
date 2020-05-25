package problem0148

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

// ListNode 是题目预定义的数据类型
type ListNode = kit.ListNode

func sortList(head *ListNode) *ListNode {
	dummyHead:=ListNode{}
	dummyHead.Next=head

	p:=head
	length:=0
	for p!=nil{
		p=p.Next
		length++
	}

	for step:=1;step<length;step=step*2{
		cur:=dummyHead.Next
		var tail *ListNode
		tail=&dummyHead
		for cur!=nil{
			left:=cur
			right:=cut(left,step)
			cur=cut(right,step)
			tail.Next=merge(left,right)
			for tail.Next!=nil{
				tail=tail.Next
			}
		}
	}
	return dummyHead.Next
}

func merge(l1 *ListNode,l2 *ListNode) *ListNode{
	dummyHead:=&ListNode{}
	new:=dummyHead
	for l1!=nil&&l2!=nil{
		if l1.Val>l2.Val {
			new.Next=l2
			l2=l2.Next
		}else{
			new.Next=l1
			l1=l1.Next
		}
		new=new.Next
	}
	if l1==nil{
		new.Next=l2
	}
	if l2==nil{
		new.Next=l1
	}
	return dummyHead.Next
}

func cut(l *ListNode,n int) *ListNode{
	p:=l
	for n>1&&p!=nil{
		p=p.Next
		n--
	}
	if p==nil{
		return nil
	}
	new:=p.Next
	p.Next=nil
	return new
}