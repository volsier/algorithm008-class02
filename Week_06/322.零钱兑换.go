/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 *
 * https://leetcode-cn.com/problems/coin-change/description/
 *
 * algorithms
 * Medium (38.83%)
 * Likes:    601
 * Dislikes: 0
 * Total Accepted:    86.7K
 * Total Submissions: 218.9K
 * Testcase Example:  '[1,2,5]\n11'
 *
 * 给定不同面额的硬币 coins 和一个总金额
 * amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。
 *
 *
 *
 * 示例 1:
 *
 * 输入: coins = [1, 2, 5], amount = 11
 * 输出: 3
 * 解释: 11 = 5 + 5 + 1
 *
 * 示例 2:
 *
 * 输入: coins = [2], amount = 3
 * 输出: -1
 *
 *
 *
 * 说明:
 * 你可以认为每种硬币的数量是无限的。
 *
 */
package main

// @lc code=start
func coinChange(coins []int, amount int) int {
	m := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		m[i] = -1
		for _, c := range coins {
			if i < c || m[i-c] == -1 {
				continue
			}
			count := m[i-c] + 1
			if m[i] == -1 || m[i] > count {
				m[i] = count
			}
		}
	}
	return m[amount]
}

// @lc code=end
