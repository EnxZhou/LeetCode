package problem0160

import "github.com/aQuaYi/LeetCode-in-Go/kit"

// ListNode is pre-defined...
type ListNode = kit.ListNode

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA==nil||headB==nil {
		return nil
	}
	pA,pB:=headA,headB
	for pA!=pB {
		if pA==nil {
			pA=headB
		}else{
			pA=pA.Next
		}
		if pB==nil{
			pB=headA
		}else{
			pB=pB.Next
		}
	}
	return pA
}
