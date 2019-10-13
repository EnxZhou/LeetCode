package problem0538

import (
	"leetcode/kit"
)

type TreeNode = kit.TreeNode

func convertBST(root *TreeNode) *TreeNode {
	a:=0
	_ = sumTree(root, &a)
	return root
}

func sumTree(root *TreeNode,sum *int) int{
	if root == nil {
		return 0
	}
	if root.Right != nil {
		root.Val+=sumTree(root.Right,sum)
	}else {
		root.Val+=*sum
	}
	*sum=root.Val
	if root.Left != nil {
		*sum=sumTree(root.Left,sum)
	}
	return *sum
}