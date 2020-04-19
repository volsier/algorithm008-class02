package main

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 *
 * https://leetcode-cn.com/problems/merge-intervals/description/
 *
 * algorithms
 * Medium (40.99%)
 * Likes:    354
 * Dislikes: 0
 * Total Accepted:    70.4K
 * Total Submissions: 171.6K
 * Testcase Example:  '[[1,3],[2,6],[8,10],[15,18]]'
 *
 * 给出一个区间的集合，请合并所有重叠的区间。
 *
 * 示例 1:
 *
 * 输入: [[1,3],[2,6],[8,10],[15,18]]
 * 输出: [[1,6],[8,10],[15,18]]
 * 解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
 *
 *
 * 示例 2:
 *
 * 输入: [[1,4],[4,5]]
 * 输出: [[1,5]]
 * 解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。
 *
 */

// @lc code=start
func merge(intervals [][]int) [][]int {
	l := len(intervals)
	if l <= 1 {
		return intervals

	}
	r := make([][]int, 0)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	t := intervals[0]
	for i := 1; i < l; i++ {
		if t[1] >= intervals[i][0] {
			if t[1] < intervals[i][1] {
				t[1] = intervals[i][1]
			}
		} else {
			r = append(r, t)
			t = intervals[i]
		}

	}
	r = append(r, t)
	return r
}

// @lc code=end
