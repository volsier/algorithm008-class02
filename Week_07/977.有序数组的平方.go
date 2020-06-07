package main

import "sort"

/*
 * @lc app=leetcode.cn id=977 lang=golang
 *
 * [977] 有序数组的平方
 *
 * https://leetcode-cn.com/problems/squares-of-a-sorted-array/description/
 *
 * algorithms
 * Easy (71.20%)
 * Likes:    82
 * Dislikes: 0
 * Total Accepted:    33.5K
 * Total Submissions: 47.2K
 * Testcase Example:  '[-4,-1,0,3,10]'
 *
 * 给定一个按非递减顺序排序的整数数组 A，返回每个数字的平方组成的新数组，要求也按非递减顺序排序。
 *
 *
 *
 * 示例 1：
 *
 * 输入：[-4,-1,0,3,10]
 * 输出：[0,1,9,16,100]
 *
 *
 * 示例 2：
 *
 * 输入：[-7,-3,2,3,11]
 * 输出：[4,9,9,49,121]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= A.length <= 10000
 * -10000 <= A[i] <= 10000
 * A 已按非递减顺序排序。
 *
 *
 */

// @lc code=start
func sortedSquares(A []int) []int {
	l := len(A)
	// B := make([]int, l)
	for i := 0; i < l; i++ {
		A[i] = A[i] * A[i]
	}
	// for i, j, k := 0, l-1, l-1; i <= j; k-- {
	// 	if A[i] > A[j] {

	// 		B[k] = A[i]
	// 		i++
	// 	} else {
	// 		B[k] = A[j]
	// 		j--
	// 	}
	// }
	sort.Ints(A)
	return A
}

// @lc code=end
