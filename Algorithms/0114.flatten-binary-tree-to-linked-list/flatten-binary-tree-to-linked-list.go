package problem0114

import (
	"leetcode/kit"
)

type TreeNode = kit.TreeNode

func flatten(root *TreeNode) {
	recur(root)
}

// recur 会 flatten root，并返回 flatten 后的 leaf 节点
func recur(root *TreeNode) *TreeNode {

}
