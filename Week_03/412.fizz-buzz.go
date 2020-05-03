package main

import "strconv"

/*
 * @lc app=leetcode.cn id=412 lang=golang
 *
 * [412] Fizz Buzz
 *
 * https://leetcode-cn.com/problems/fizz-buzz/description/
 *
 * algorithms
 * Easy (63.18%)
 * Likes:    56
 * Dislikes: 0
 * Total Accepted:    32.4K
 * Total Submissions: 51.2K
 * Testcase Example:  '1'
 *
 * 写一个程序，输出从 1 到 n 数字的字符串表示。
 *
 * 1. 如果 n 是3的倍数，输出“Fizz”；
 *
 * 2. 如果 n 是5的倍数，输出“Buzz”；
 *
 * 3.如果 n 同时是3和5的倍数，输出 “FizzBuzz”。
 *
 * 示例：
 *
 * n = 15,
 *
 * 返回:
 * [
 * ⁠   "1",
 * ⁠   "2",
 * ⁠   "Fizz",
 * ⁠   "4",
 * ⁠   "Buzz",
 * ⁠   "Fizz",
 * ⁠   "7",
 * ⁠   "8",
 * ⁠   "Fizz",
 * ⁠   "Buzz",
 * ⁠   "11",
 * ⁠   "Fizz",
 * ⁠   "13",
 * ⁠   "14",
 * ⁠   "FizzBuzz"
 * ]
 *
 *
 */

// @lc code=start
func fizzBuzz(n int) []string {
	r := make([]string, n+1)
	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			r[i-1] += "Fizz"
		}
		if i%5 == 0 {
			r[i-1] += "Buzz"
		}
		if r[i-1] == "" {
			r[i-1] = strconv.Itoa(i)
		}
	}
	return r
}

// @lc code=end
