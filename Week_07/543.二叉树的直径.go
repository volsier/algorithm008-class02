package main

/*
 * @lc app=leetcode.cn id=543 lang=golang
 *
 * [543] 二叉树的直径
 *
 * https://leetcode-cn.com/problems/diameter-of-binary-tree/description/
 *
 * algorithms
 * Easy (49.33%)
 * Likes:    325
 * Dislikes: 0
 * Total Accepted:    45.7K
 * Total Submissions: 92K
 * Testcase Example:  '[1,2,3,4,5]'
 *
 * 给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过也可能不穿过根结点。
 *
 *
 *
 * 示例 :
 * 给定二叉树
 *
 * ⁠         1
 * ⁠        / \
 * ⁠       2   3
 * ⁠      / \
 * ⁠     4   5
 *
 *
 * 返回 3, 它的长度是路径 [4,2,1,3] 或者 [5,2,1,3]。
 *
 *
 *
 * 注意：两结点之间的路径长度是以它们之间边的数目表示。
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
func diameterOfBinaryTree(root *TreeNode) int {
	m := 0
	dep(root, &m)
	return m
}
func dep(node *TreeNode, m *int) int {
	if node == nil {
		return 0
	}
	l := dep(node.Left, m)
	r := dep(node.Right, m)
	p := max(l, r)
	*m = max(l+r, *m)
	return p + 1
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
