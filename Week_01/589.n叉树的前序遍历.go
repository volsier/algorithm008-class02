package main

/*
 * @lc app=leetcode.cn id=589 lang=golang
 *
 * [589] N叉树的前序遍历
 *
 * https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal/description/
 *
 * algorithms
 * Easy (72.75%)
 * Likes:    75
 * Dislikes: 0
 * Total Accepted:    23K
 * Total Submissions: 31.7K
 * Testcase Example:  '[1,null,3,2,4,null,5,6]'
 *
 * 给定一个 N 叉树，返回其节点值的前序遍历。
 *
 * 例如，给定一个 3叉树 :
 *
 *
 *
 *
 *
 *
 *
 * 返回其前序遍历: [1,3,5,6,2,4]。
 *
 *
 *
 * 说明: 递归法很简单，你可以使用迭代法完成此题吗?
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
 */
type Stack struct {
	Value *Node
	Next  *Stack
}

func preorder(root *Node) []int {
	s := make([]int, 0)
	if root == nil {
		return s
	}
	stack := &Stack{
		Value: root,
	}
	for v := stack; v != nil; v = v.Next {
		s = append(s, v.Value.Val)
		n := v.Next
		c := v
		for _, vv := range v.Value.Children {
			c.Next = &Stack{
				Value: vv,
			}
			c = c.Next
		}
		c.Next = n
	}
	return s
}

// @lc code=end
