/*
 * @lc app=leetcode.cn id=433 lang=golang
 *
 * [433] 最小基因变化
 *
 * https://leetcode-cn.com/problems/minimum-genetic-mutation/description/
 *
 * algorithms
 * Medium (49.25%)
 * Likes:    32
 * Dislikes: 0
 * Total Accepted:    4.2K
 * Total Submissions: 8.3K
 * Testcase Example:  '"AACCGGTT"\n"AACCGGTA"\n["AACCGGTA"]'
 *
 * 一条基因序列由一个带有8个字符的字符串表示，其中每个字符都属于 "A", "C", "G", "T"中的任意一个。
 *
 * 假设我们要调查一个基因序列的变化。一次基因变化意味着这个基因序列中的一个字符发生了变化。
 *
 * 例如，基因序列由"AACCGGTT" 变化至 "AACCGGTA" 即发生了一次基因变化。
 *
 * 与此同时，每一次基因变化的结果，都需要是一个合法的基因串，即该结果属于一个基因库。
 *
 * 现在给定3个参数 — start, end,
 * bank，分别代表起始基因序列，目标基因序列及基因库，请找出能够使起始基因序列变化为目标基因序列所需的最少变化次数。如果无法实现目标变化，请返回
 * -1。
 *
 * 注意:
 *
 *
 * 起始基因序列默认是合法的，但是它并不一定会出现在基因库中。
 * 所有的目标基因序列必须是合法的。
 * 假定起始基因序列与目标基因序列是不一样的。
 *
 *
 * 示例 1:
 *
 *
 * start: "AACCGGTT"
 * end:   "AACCGGTA"
 * bank: ["AACCGGTA"]
 *
 * 返回值: 1
 *
 *
 * 示例 2:
 *
 *
 * start: "AACCGGTT"
 * end:   "AAACGGTA"
 * bank: ["AACCGGTA", "AACCGCTA", "AAACGGTA"]
 *
 * 返回值: 2
 *
 *
 * 示例 3:
 *
 *
 * start: "AAAAACCC"
 * end:   "AACCCCCC"
 * bank: ["AAAACCCC", "AAACCCCC", "AACCCCCC"]
 *
 * 返回值: 3
 *
 *
 */
package main

// @lc code=start
func minMutation(start string, end string, bank []string) int {
	l := len(bank)
	if l == 0 {
		return -1
	}
	visted := make([]int, l)
	queue := []string{start}
	n := 0
	for len(queue) > 0 {
		n++
		qL := len(queue)
		for i := 0; i < qL; i++ {

			for j := 0; j < l; j++ {
				if visted[j] == 0 && bank[j] != start {
					dif := 0
					for s := 0; s < 8; s++ {
						if bank[j][s] != queue[i][s] {
							dif++
						}
						if dif > 1 {
							break
						}
					}
					if dif == 1 {
						visted[j] = n
						if bank[j] == end {
							return n
						}
						queue = append(queue, bank[j])
					}
				}
			}
		}
		queue = queue[qL:]
	}
	return -1
}

// @lc code=end
