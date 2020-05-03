package main

import "math"

/*
 * @lc app=leetcode.cn id=111 lang=golang
 *
 * [111] 二叉树的最小深度
 *
 * https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/description/
 *
 * algorithms
 * Easy (41.97%)
 * Likes:    245
 * Dislikes: 0
 * Total Accepted:    68.8K
 * Total Submissions: 163.4K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，找出其最小深度。
 *
 * 最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
 *
 * 说明: 叶子节点是指没有子节点的节点。
 *
 * 示例:
 *
 * 给定二叉树 [3,9,20,null,null,15,7],
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 * 返回它的最小深度  2.
 *
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	//递归
	md := math.MaxInt64
	checkminDepth(root, 1, &md)
	return md
}
func checkminDepth(root *TreeNode, l int, md *int) {
	if root.Left == nil && root.Right == nil {
		if l < *md {
			*md = l
		}
	}
	if root.Left != nil {
		checkminDepth(root.Left, l+1, md)
	}
	if root.Right != nil {
		checkminDepth(root.Right, l+1, md)
	}
}

// @lc code=end
