package problem0114

import (
	"leetcode/kit"
)

type TreeNode = kit.TreeNode

func flatten(root *TreeNode) {
	root=recur(root)
}

// recur 会 flatten root，并返回 flatten 后的 leaf 节点
func recur(root *TreeNode) *TreeNode {
	if root==nil{
		return nil
	}
	if root.Left==nil&&root.Right==nil{
		return root
	}
	root.Left=recur(root.Left)
	root.Right=recur(root.Right)
	tmp:=root.Left
	if tmp!=nil {
		for tmp.Right!=nil {
			tmp=tmp.Right
		}
		tmp.Right=root.Right
		root.Right=root.Left
		root.Left=nil
	}
	return root
}

func flatten1(root *TreeNode) {
	for root!=nil {
		if root.Left!=nil {
			tmp:=root.Left
			for tmp.Right!=nil {
				tmp=tmp.Right
			}
			tmp.Right=root.Right
			root.Right=root.Left
			root.Left=nil
		}else{
			root = root.Right
		}
	}
}