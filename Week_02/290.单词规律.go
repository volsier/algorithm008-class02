package main

import "strings"

/*
 * @lc app=leetcode.cn id=290 lang=golang
 *
 * [290] 单词规律
 *
 * https://leetcode-cn.com/problems/word-pattern/description/
 *
 * algorithms
 * Easy (42.48%)
 * Likes:    146
 * Dislikes: 0
 * Total Accepted:    21.5K
 * Total Submissions: 50.7K
 * Testcase Example:  '"abba"\n"dog cat cat dog"'
 *
 * 给定一种规律 pattern 和一个字符串 str ，判断 str 是否遵循相同的规律。
 *
 * 这里的 遵循 指完全匹配，例如， pattern 里的每个字母和字符串 str 中的每个非空单词之间存在着双向连接的对应规律。
 *
 * 示例1:
 *
 * 输入: pattern = "abba", str = "dog cat cat dog"
 * 输出: true
 *
 * 示例 2:
 *
 * 输入:pattern = "abba", str = "dog cat cat fish"
 * 输出: false
 *
 * 示例 3:
 *
 * 输入: pattern = "aaaa", str = "dog cat cat dog"
 * 输出: false
 *
 * 示例 4:
 *
 * 输入: pattern = "abba", str = "dog dog dog dog"
 * 输出: false
 *
 * 说明:
 * 你可以假设 pattern 只包含小写字母， str 包含了由单个空格分隔的小写字母。
 *
 */

// @lc code=start
func wordPattern(pattern string, str string) bool {
	p := strings.Split(pattern, "")
	s := strings.Split(str, " ")
	if len(p) != len(s) {
		return false
	}
	pL := len(p)
	mP := make(map[string]string)
	mS := make(map[string]string)
	for i := 0; i < pL; i++ {
		pv, pOk := mP[p[i]]
		sv, sOk := mS[s[i]]
		if pOk || sOk {
			if pv != s[i] || sv != p[i] {
				return false
			}
		} else {
			mP[p[i]] = s[i]
			mS[s[i]] = p[i]
		}
	}
	return true
}

// @lc code=end
