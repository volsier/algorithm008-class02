package main

import (
	"math"
	"sort"
)

/*
 * @lc app=leetcode.cn id=1200 lang=golang
 *
 * [1200] 最小绝对差
 *
 * https://leetcode-cn.com/problems/minimum-absolute-difference/description/
 *
 * algorithms
 * Easy (65.87%)
 * Likes:    18
 * Dislikes: 0
 * Total Accepted:    7.6K
 * Total Submissions: 11.4K
 * Testcase Example:  '[4,2,1,3]'
 *
 * 给你个整数数组 arr，其中每个元素都 不相同。
 *
 * 请你找到所有具有最小绝对差的元素对，并且按升序的顺序返回。
 *
 *
 *
 * 示例 1：
 *
 * 输入：arr = [4,2,1,3]
 * 输出：[[1,2],[2,3],[3,4]]
 *
 *
 * 示例 2：
 *
 * 输入：arr = [1,3,6,10,15]
 * 输出：[[1,3]]
 *
 *
 * 示例 3：
 *
 * 输入：arr = [3,8,-10,23,19,-4,-14,27]
 * 输出：[[-14,-10],[19,23],[23,27]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 2 <= arr.length <= 10^5
 * -10^6 <= arr[i] <= 10^6
 *
 *
 */

// @lc code=start
func minimumAbsDifference(arr []int) [][]int {
	l := len(arr)
	r := make([][]int, 0)
	sort.Ints(arr)
	min := math.MaxInt64
	for i := 1; i < l; i++ {
		min = cMin(arr[i]-arr[i-1], min)
	}
	for i := 1; i < l; i++ {
		if arr[i]-arr[i-1] == min {
			r = append(r, []int{arr[i-1], arr[i]})
		}
	}

	return r
}
func cMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
