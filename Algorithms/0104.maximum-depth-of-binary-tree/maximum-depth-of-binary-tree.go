package problem0104

import (
	"leetcode/kit"
)

type TreeNode = kit.TreeNode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}else{
		return max(maxDepth(root.Right),maxDepth(root.Left))+1
	}
}

func maxTree(root *TreeNode,num int) int {
	if root == nil {
		return num
	}
	right:=maxTree(root.Right,num+1)
	left:=maxTree(root.Left,num+1)
	if right>left {
		return right
	}else {
		return left
	}
}

func max(a,b int)int{
	if a > b {
		return a
	}else{
		return b
	}
}