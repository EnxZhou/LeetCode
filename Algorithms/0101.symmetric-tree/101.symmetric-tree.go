/*
 * @lc app=leetcode id=101 lang=golang
 *
 * [101] Symmetric Tree
 *
 * https://leetcode.com/problems/symmetric-tree/description/
 *
 * algorithms
 * Easy (44.47%)
 * Likes:    2652
 * Dislikes: 57
 * Total Accepted:    465.2K
 * Total Submissions: 1M
 * Testcase Example:  '[1,2,2,3,4,4,3]'
 *
 * Given a binary tree, check whether it is a mirror of itself (ie, symmetric
 * around its center).
 *
 * For example, this binary tree [1,2,2,3,4,4,3] is symmetric:
 *
 *
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   2
 * ⁠/ \ / \
 * 3  4 4  3
 *
 *
 *
 *
 * But the following [1,2,2,null,3,null,3] is not:
 *
 *
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   2
 * ⁠  \   \
 * ⁠  3    3
 *
 *
 *
 *
 * Note:
 * Bonus points if you could solve it both recursively and iteratively.
 *
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package problem0101

import (
	"leetcode/kit"
)

type TreeNode = kit.TreeNode

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}
	isMirror(root.Left, root.Right)
	return false
}
func isMirror(a *TreeNode, b *TreeNode) bool {
	if a != nil {

	}
	if a.Val != b.Val {
		return false
	}
	if isMirror(a.Left, b.Right) && isMirror(a.Right, b.Left) {
		return true
	} else {
		return false
	}
}
