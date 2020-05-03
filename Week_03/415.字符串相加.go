/*
 * @lc app=leetcode.cn id=415 lang=golang
 *
 * [415] 字符串相加
 *
 * https://leetcode-cn.com/problems/add-strings/description/
 *
 * algorithms
 * Easy (49.78%)
 * Likes:    155
 * Dislikes: 0
 * Total Accepted:    31.4K
 * Total Submissions: 63.1K
 * Testcase Example:  '"0"\n"0"'
 *
 * 给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。
 *
 * 注意：
 *
 *
 * num1 和num2 的长度都小于 5100.
 * num1 和num2 都只包含数字 0-9.
 * num1 和num2 都不包含任何前导零。
 * 你不能使用任何內建 BigInteger 库， 也不能直接将输入的字符串转换为整数形式。
 *
 *
 */
package main

// @lc code=start
func addStrings(num1 string, num2 string) string {
	l1 := len(num1)
	l2 := len(num2)
	var r []byte
	if l1 > l2 {
		r = make([]byte, l1+1)
	} else {
		r = make([]byte, l2+1)
	}
	ad := 0
	s := len(r) - 1
	for i, j := l1-1, l2-1; (i >= 0 || j >= 0) && s > 0; {
		sum := ad
		if i >= 0 {
			sum += int(num1[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(num2[j] - '0')
			j--
		}
		ad = sum / 10
		r[s] = byte(sum%10 + '0')
		s--

	}
	if ad == 1 {
		r[s] = '1'
		s--
	}
	return string(r[s+1:])
}

// @lc code=end
