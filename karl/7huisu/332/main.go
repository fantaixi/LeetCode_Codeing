package _32

import "sort"

/*
给定一个机票的字符串二维数组 [from, to]，子数组中的两个成员分别表示飞机出发和降落的机场地点，
对该行程进行重新规划排序。所有这些机票都属于一个从 JFK（肯尼迪国际机场）出发的先生，所以该行程必须从 JFK 开始。
提示：
如果存在多种有效的行程，请你按字符自然排序返回最小的行程组合。例如，行程 ["JFK", "LGA"] 与 ["JFK", "LGB"] 相比就更小，排序更靠前
所有的机场都用三个大写字母表示（机场代码）。
假定所有机票至少存在一种合理的行程。
所有的机票必须都用一次 且 只能用一次。
 */
/*
已知图中存在欧拉路径，你要找出一个欧拉路径，可以用 hierholzer 算法。
任选一个点为起点（题目告诉你了），遍历它所有邻接的边（设置不同的分支）。
DFS 搜索，访问邻接的点，并且将走过的边（邻接关系）删除。
如果走到的当前点，已经没有相邻边了，则将当前点推入 res。
随着递归的出栈，点不断推入 res 的开头，最后就得到一个从起点出发的欧拉路径。
 */
/*
方便对数据进行排序，而且还要容易增删元素
 */
type pair struct {
	target string
	visited bool
}
type pairs []*pair

func (p pairs) Len() int {
	return len(p)
}
func (p pairs) Swap(i, j int) {
	p[i],p[j] = p[j],p[i]
}
func (p pairs) Less(i, j int) bool {
	return p[i].target < p[j].target
}
func findItinerary(tickets [][]string) []string {
	res := []string{}
	// map[出发机场] pair{目的地,是否被访问过}
	targets := make(map[string]pairs)
	for _, ticket := range tickets {
		if targets[ticket[0]] == nil{
			targets[ticket[0]] = make(pairs,0)
		}
		targets[ticket[0]] = append(targets[ticket[0]],&pair{target: ticket[1], visited: false})
	}
	for k, _ := range targets {
		sort.Sort(targets[k])
	}
	res = append(res,"JFK")
	var backtrack func() bool
	backtrack = func() bool {
		//回溯遍历的过程中，遇到的机场个数，如果达到了（航班数量+1），
		//那么我们就找到了一个行程，把所有航班串在一起了
		if len(tickets)+1 == len(res) {
			return true
		}
		// 取出起飞航班对应的目的地
		for _, p := range targets[res[len(res)-1]] {
			if p.visited == false {
				res = append(res,p.target)
				p.visited = true
				if backtrack() {
					return true
				}
				res = res[:len(res)-1]
				p.visited = false
			}
		}
		return false
	}
	backtrack()
	return res
}