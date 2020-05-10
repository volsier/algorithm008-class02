/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 *
 * https://leetcode-cn.com/problems/number-of-islands/description/
 *
 * algorithms
 * Medium (47.97%)
 * Likes:    548
 * Dislikes: 0
 * Total Accepted:    101.3K
 * Total Submissions: 205.4K
 * Testcase Example:  '[["1","1","1","1","0"],["1","1","0","1","0"],["1","1","0","0","0"],["0","0","0","0","0"]]'
 *
 * 给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
 *
 * 岛屿总是被水包围，并且每座岛屿只能由水平方向或竖直方向上相邻的陆地连接形成。
 *
 * 此外，你可以假设该网格的四条边均被水包围。
 *
 *
 *
 * 示例 1:
 *
 * 输入:
 * 11110
 * 11010
 * 11000
 * 00000
 * 输出: 1
 *
 *
 * 示例 2:
 *
 * 输入:
 * 11000
 * 11000
 * 00100
 * 00011
 * 输出: 3
 * 解释: 每座岛屿只能由水平和/或竖直方向上相邻的陆地连接而成。
 *
 *
 */
package main

// @lc code=start
func numIslands(grid [][]byte) int {
	iL := len(grid)
	if iL == 0 {
		return 0
	}
	jL := len(grid[0])
	r := 0
	for i := 0; i < iL; i++ {
		for j := 0; j < jL; j++ {
			if grid[i][j] == '1' {
				dF(grid, i, j)
				r++
			}
		}
	}
	return r
}
func dF(grid [][]byte, i, j int) {
	grid[i][j] = '0'
	if i-1 >= 0 && grid[i-1][j] == '1' {
		dF(grid, i-1, j)
	}
	if i+1 < len(grid) && grid[i+1][j] == '1' {
		dF(grid, i+1, j)
	}
	if j-1 >= 0 && grid[i][j-1] == '1' {
		dF(grid, i, j-1)
	}
	if j+1 < len(grid[0]) && grid[i][j+1] == '1' {
		dF(grid, i, j+1)
	}
}

// @lc code=end
