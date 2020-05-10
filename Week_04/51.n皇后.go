/*
 * @lc app=leetcode.cn id=51 lang=golang
 *
 * [51] N皇后
 *
 * https://leetcode-cn.com/problems/n-queens/description/
 *
 * algorithms
 * Hard (68.67%)
 * Likes:    396
 * Dislikes: 0
 * Total Accepted:    39K
 * Total Submissions: 56.2K
 * Testcase Example:  '4'
 *
 * n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
 *
 *
 *
 * 上图为 8 皇后问题的一种解法。
 *
 * 给定一个整数 n，返回所有不同的 n 皇后问题的解决方案。
 *
 * 每一种解法包含一个明确的 n 皇后问题的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
 *
 * 示例:
 *
 * 输入: 4
 * 输出: [
 * ⁠[".Q..",  // 解法 1
 * ⁠ "...Q",
 * ⁠ "Q...",
 * ⁠ "..Q."],
 *
 * ⁠["..Q.",  // 解法 2
 * ⁠ "Q...",
 * ⁠ "...Q",
 * ⁠ ".Q.."]
 * ]
 * 解释: 4 皇后问题存在两个不同的解法。
 *
 *
 *
 *
 * 提示：
 *
 *
 *
 * 皇后，是国际象棋中的棋子，意味着国王的妻子。皇后只做一件事，那就是“吃子”。当她遇见可以吃的棋子时，就迅速冲上去吃掉棋子。当然，她横、竖、斜都可走一或七步，可进可退。（引用自
 * 百度百科 - 皇后 ）
 *
 *
 */
package main

// @lc code=start
func solveNQueens(n int) [][]string {
	r := make([][]string, 0)
	if n == 0 {
		return r
	}
	cols := make([]bool, n)
	pies := make([]bool, 2*n-1)
	nas := make([]bool, 2*n-1)
	res := make([][]int, 0)
	dsf([]int{}, cols, pies, nas, &res, n)
	r = printRes(res, n)
	return r
}
func dsf(rows []int, cols []bool, pies []bool, nas []bool, res *[][]int, n int) {
	row := len(rows)
	if row == n {
		tmp := make([]int, n)
		copy(tmp, rows)
		*res = append(*res, tmp)
		return
	}
	for col := 0; col < n; col++ {
		na := n - 1 - (row - col)
		pie := row + col
		if cols[col] || pies[pie] || nas[na] {
			continue
		}
		cols[col] = true
		pies[pie] = true
		nas[na] = true
		dsf(append(rows, col), cols, pies, nas, res, n)

		cols[col] = false
		pies[pie] = false
		nas[na] = false
	}
}
func printRes(res [][]int, n int) [][]string {
	result := make([][]string, 0)
	for _, v := range res {
		var s []string
		for _, val := range v {
			str := ""
			for i := 0; i < n; i++ {
				if i == val {
					str += "Q"
				} else {
					str += "."
				}
			}
			s = append(s, str)
		}
		result = append(result, s)
	}
	return result

}

// @lc code=end
