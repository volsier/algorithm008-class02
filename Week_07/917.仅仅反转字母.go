package main

import (
	"fmt"
	"sort"
	"unicode"
)

/*
 * @lc app=leetcode.cn id=917 lang=golang
 *
 * [917] 仅仅反转字母
 *
 * https://leetcode-cn.com/problems/reverse-only-letters/description/
 *
 * algorithms
 * Easy (54.09%)
 * Likes:    44
 * Dislikes: 0
 * Total Accepted:    10.7K
 * Total Submissions: 19.8K
 * Testcase Example:  '"ab-cd"'
 *
 * 给定一个字符串 S，返回 “反转后的” 字符串，其中不是字母的字符都保留在原地，而所有字母的位置发生反转。
 *
 *
 *
 *
 *
 *
 * 示例 1：
 *
 * 输入："ab-cd"
 * 输出："dc-ba"
 *
 *
 * 示例 2：
 *
 * 输入："a-bC-dEf-ghIj"
 * 输出："j-Ih-gfE-dCba"
 *
 *
 * 示例 3：
 *
 * 输入："Test1ng-Leet=code-Q!"
 * 输出："Qedo1ct-eeLg=ntse-T!"
 *
 *
 *
 *
 * 提示：
 *
 *
 * S.length <= 100
 * 33 <= S[i].ASCIIcode <= 122
 * S 中不包含 \ or "
 *
 *
 */

// @lc code=start
func reverseOnlyLetters(S string) string {
	l := len(S)
	C := []rune(S)
	j := l - 1
	sort.Slice()

	strings.
	for i := 0; i < j; {
		if unicode.IsLetter(C[i]) && unicode.IsLetter(C[j]) {
			C[i], C[j] = C[j], C[i]
			i++
		} else {
			if !unicode.IsLetter(C[i]) {
				i++
				continue
			}
		}
		j--
	}
	return string(C)
}

// @lc code=end
