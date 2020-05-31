/*
 * @lc app=leetcode.cn id=312 lang=golang
 *
 * [312] 戳气球
 *
 * https://leetcode-cn.com/problems/burst-balloons/description/
 *
 * algorithms
 * Hard (58.67%)
 * Likes:    294
 * Dislikes: 0
 * Total Accepted:    12.5K
 * Total Submissions: 20.8K
 * Testcase Example:  '[3,1,5,8]'
 *
 * 有 n 个气球，编号为0 到 n-1，每个气球上都标有一个数字，这些数字存在数组 nums 中。
 *
 * 现在要求你戳破所有的气球。每当你戳破一个气球 i 时，你可以获得 nums[left] * nums[i] * nums[right] 个硬币。 这里的
 * left 和 right 代表和 i 相邻的两个气球的序号。注意当你戳破了气球 i 后，气球 left 和气球 right 就变成了相邻的气球。
 *
 * 求所能获得硬币的最大数量。
 *
 * 说明:
 *
 *
 * 你可以假设 nums[-1] = nums[n] = 1，但注意它们不是真实存在的所以并不能被戳破。
 * 0 ≤ n ≤ 500, 0 ≤ nums[i] ≤ 100
 *
 *
 * 示例:
 *
 * 输入: [3,1,5,8]
 * 输出: 167
 * 解释: nums = [3,1,5,8] --> [3,5,8] -->   [3,8]   -->  [8]  --> []
 * coins =  3*1*5      +  3*5*8    +  1*3*8      + 1*8*1   = 167
 *
 *
 */
package main

// @lc code=start
func maxCoins(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	nums = append(nums, 1)
	nums = append([]int{1}, nums...)
	n := len(nums)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[i][i] = 0
		if i > 0 {
			dp[i-1][i] = 0
		}
		if i > 1 {
			dp[i-2][i] = nums[i-1] * nums[i-2] * nums[i]
		}
	}
	for i := 3; i < n; i++ {
		for j := i - 3; j >= 0; j-- {
			for k := j + 1; k < i; k++ {
				dp[j][i] = max(dp[j][i], dp[j][k]+dp[k][i]+nums[k]*nums[i]*nums[j])
			}
		}
	}
	return dp[0][n-1]

}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
