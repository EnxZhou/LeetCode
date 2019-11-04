package problem0226

import (
	"leetcode/kit"
)

type TreeNode = kit.TreeNode

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	cache:=invertTree(root.Right)
	root.Right=invertTree(root.Left)
	root.Left=cache
	return root
}
