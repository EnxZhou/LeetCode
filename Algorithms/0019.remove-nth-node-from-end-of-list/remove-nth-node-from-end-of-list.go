package problem0019

/*
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// ListNode 是节点
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	dummyHead:=&ListNode{}
	dummyHead.Next=head
	nodeMaps:=make([]*ListNode,0)
	for head!=nil{
		oneMap:=head
		nodeMaps=append(nodeMaps,oneMap)
		head=head.Next
	}
	if len(nodeMaps)==n{
		return dummyHead.Next.Next
	}
	deletePreNode:=nodeMaps[len(nodeMaps)-n-1]
	deletePreNode.Next=deletePreNode.Next.Next
	return dummyHead.Next
}

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	dummyHead:=&ListNode{}
	dummyHead.Next=head
	nodeMaps:=make(map[int]*ListNode,0)
	for i:=0;head!=nil;i++{
		nodeMaps[i]=head
		head=head.Next
	}
	if len(nodeMaps)==n{
		return dummyHead.Next.Next
	}
	deletePreNode:=nodeMaps[len(nodeMaps)-n-1]
	deletePreNode.Next=deletePreNode.Next.Next
	return dummyHead.Next
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fast,slow:=head,head
	for n>0{
		fast=fast.Next
		if fast==nil{
			return head.Next
		}
		n--
	}
	for fast.Next!=nil{
		fast=fast.Next
		slow=slow.Next
	}
	slow.Next=slow.Next.Next
	return head
}