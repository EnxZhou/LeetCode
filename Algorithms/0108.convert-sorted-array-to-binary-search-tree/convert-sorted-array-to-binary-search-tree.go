package problem0108

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type TreeNode = kit.TreeNode

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums)==0 {
		return nil
	}
	var middle int
	middle=len(nums)/2
	root:=&TreeNode{
		Val:nums[middle],
		Left: sortedArrayToBST(nums[:middle]),
		Right: sortedArrayToBST(nums[middle+1:]),
	}
	return root
}
