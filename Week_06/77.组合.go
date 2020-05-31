/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 *
 * https://leetcode-cn.com/problems/combinations/description/
 *
 * algorithms
 * Medium (73.41%)
 * Likes:    273
 * Dislikes: 0
 * Total Accepted:    51.9K
 * Total Submissions: 70.4K
 * Testcase Example:  '4\n2'
 *
 * 给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
 *
 * 示例:
 *
 * 输入: n = 4, k = 2
 * 输出:
 * [
 * ⁠ [2,4],
 * ⁠ [3,4],
 * ⁠ [2,3],
 * ⁠ [1,2],
 * ⁠ [1,3],
 * ⁠ [1,4],
 * ]
 *
 */
package main

// @lc code=start
func combine(n int, k int) [][]int {
	check := make([]int, n)
	for i := 1; i <= n; i++ {
		check[i-1] = i
	}
	result := make([][]int, 0)
	brack([]int{}, check, k, &result)
	return result
}
func brack(cur []int, check []int, k int, result *[][]int) {
	if len(cur) == k {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*result = append(*result, tmp)
	}
	for i := 0; i < len(check); i++ {
		brack(append(cur, check[i]), check[i+1:], k, result)
	}
}

// @lc code=end
