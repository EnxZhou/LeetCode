package problem0021

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

func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	fakeNew:=&ListNode{}
	new:=fakeNew
	for l1!=nil&&l2!=nil {
		if l1.Val> l2.Val {
			new.Next=l2
			l2=l2.Next
		}else{
			new.Next=l1
			l1=l1.Next
		}
		new=new.Next
	}
	if l1==nil {
		new.Next=l2
	}
	if l2==nil {
		new.Next=l1
	}
	return fakeNew.Next
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1==nil {
		return l2
	}
	if l2==nil {
		return l1
	}
	if l1.Val>l2.Val {
		l2.Next=mergeTwoLists(l1,l2.Next)
		return l2
	}
	l1.Next=mergeTwoLists(l1.Next,l2)
	return l1
}