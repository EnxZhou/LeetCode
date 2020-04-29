package problem0337

import (
	"leetcode/kit"
)

type TreeNode = kit.TreeNode

func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	sum1, sum2 := root.Val, 0
	if root.Left != nil {
		sum1 += rob(root.Left.Left) + rob(root.Left.Right)
	}
	if root.Right != nil {
		sum1 += rob(root.Right.Left) + rob(root.Right.Right)
	}
	sum2 = rob(root.Left) + rob(root.Right)
	return max(sum1, sum2)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
