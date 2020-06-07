package main

/*
 * @lc app=leetcode.cn id=542 lang=golang
 *
 * [542] 01 矩阵
 *
 * https://leetcode-cn.com/problems/01-matrix/description/
 *
 * algorithms
 * Medium (38.32%)
 * Likes:    237
 * Dislikes: 0
 * Total Accepted:    26.8K
 * Total Submissions: 61.3K
 * Testcase Example:  '[[0,0,0],[0,1,0],[0,0,0]]'
 *
 * 给定一个由 0 和 1 组成的矩阵，找出每个元素到最近的 0 的距离。
 *
 * 两个相邻元素间的距离为 1 。
 *
 * 示例 1:
 * 输入:
 *
 *
 * 0 0 0
 * 0 1 0
 * 0 0 0
 *
 *
 * 输出:
 *
 *
 * 0 0 0
 * 0 1 0
 * 0 0 0
 *
 *
 * 示例 2:
 * 输入:
 *
 *
 * 0 0 0
 * 0 1 0
 * 1 1 1
 *
 *
 * 输出:
 *
 *
 * 0 0 0
 * 0 1 0
 * 1 2 1
 *
 *
 * 注意:
 *
 *
 * 给定矩阵的元素个数不超过 10000。
 * 给定矩阵中至少有一个元素是 0。
 * 矩阵中的元素只在四个方向上相邻: 上、下、左、右。
 *
 *
 */

// @lc code=start
func updateMatrix(matrix [][]int) [][]int {
	l := len(matrix)
	if l == 0 {
		return matrix
	}
	jL := len(matrix[0])
	ra := jL * l
	for i := 0; i < l; i++ {
		for j := 0; j < jL; j++ {
			if matrix[i][j] != 0 {
				lCell := ra
				uCell := ra
				if i-1 >= 0 {
					lCell = matrix[i-1][j]
				}
				if j-1 >= 0 {
					uCell = matrix[i][j-1]
				}
				matrix[i][j] = checkMin(lCell, uCell) + 1

			}
		}
	}
	for i := l - 1; i >= 0; i-- {
		for j := jL - 1; j >= 0; j-- {
			if matrix[i][j] != 0 {
				rCell := ra
				dCell := ra
				if i+1 < l {
					rCell = matrix[i+1][j]
				}
				if j+1 < jL {
					dCell = matrix[i][j+1]
				}
				matrix[i][j] = checkMin(checkMin(rCell, dCell)+1, matrix[i][j])

			}
		}
	}
	return matrix
}
func checkMin(nua int, nub int) int {
	if nua > nub {
		return nub
	}
	return nua
}

// @lc code=end
