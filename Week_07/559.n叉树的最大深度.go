package main

/*
 * @lc app=leetcode.cn id=559 lang=golang
 *
 * [559] N叉树的最大深度
 *
 * https://leetcode-cn.com/problems/maximum-depth-of-n-ary-tree/description/
 *
 * algorithms
 * Easy (68.94%)
 * Likes:    84
 * Dislikes: 0
 * Total Accepted:    20.9K
 * Total Submissions: 30.2K
 * Testcase Example:  '[1,null,3,2,4,null,5,6]'
 *
 * 给定一个 N 叉树，找到其最大深度。
 *
 * 最大深度是指从根节点到最远叶子节点的最长路径上的节点总数。
 *
 * 例如，给定一个 3叉树 :
 *
 *
 *
 *
 *
 *
 *
 * 我们应返回其最大深度，3。
 *
 * 说明:
 *
 *
 * 树的深度不会超过 1000。
 * 树的节点总不会超过 5000。
 *
 */
type Node struct {
	Val      int
	Children []*Node
}

// @lc code=start
/**
* Definition for a Node.
* type Node struct {
*     Val int
*     Children []*Node
* }
[1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
*/

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}

	l := 1
	l = checkMaxDepth(root, l)
	return l
}
func checkMaxDepth(node *Node, l int) int {
	if len(node.Children) != 0 {
		l++
		tL := l
		for _, v := range node.Children {
			tL = max(tL, checkMaxDepth(v, l))
		}
		l = max(l, tL)
	}

	return l
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
