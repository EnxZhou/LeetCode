package problem0437

import (
	"leetcode/kit"
)

type TreeNode = kit.TreeNode

func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	count:=0

	count+=dfs(root,sum)
	count+=pathSum(root.Left,sum)
	count+=pathSum(root.Right,sum)
	return count
}

func dfs(root *TreeNode,sum int) int {
	if root == nil {
		return 0
	}
	count:=0
	if root.Val == sum {
		count++
	}
	count+=dfs(root.Left,sum-root.Val)
	count+=dfs(root.Right,sum-root.Val)
	return count
}