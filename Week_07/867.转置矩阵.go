package main

/*
 * @lc app=leetcode.cn id=867 lang=golang
 *
 * [867] 转置矩阵
 *
 * https://leetcode-cn.com/problems/transpose-matrix/description/
 *
 * algorithms
 * Easy (67.26%)
 * Likes:    88
 * Dislikes: 0
 * Total Accepted:    18K
 * Total Submissions: 26.7K
 * Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]'
 *
 * 给定一个矩阵 A， 返回 A 的转置矩阵。
 *
 * 矩阵的转置是指将矩阵的主对角线翻转，交换矩阵的行索引与列索引。
 *
 *
 *
 * 示例 1：
 *
 * 输入：[[1,2,3],[4,5,6],[7,8,9]]
 * 输出：[[1,4,7],[2,5,8],[3,6,9]]
 * 123
 123
 456
 789

 147
 258
 369
 14
 25
 36
 *
 * 示例 2：
 *
 * 输入：[[1,2,3],[4,5,6]]
 * 输出：[[1,4],[2,5],[3,6]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= A.length <= 1000
 * 1 <= A[0].length <= 1000
 *
 *
*/

// @lc code=start
func transpose(A [][]int) [][]int {
	l := len(A)
	xl := len(A[0])
	B := make([][]int, xl)
	for i := 0; i < xl; i++ {
		B[i] = make([]int, l)
		for j := 0; j < l; j++ {
			B[i][j] = A[j][i]
		}
	}
	return B
}

// @lc code=end
