package main

func main() {

}
/*
给定一个 N 叉树，找到其最大深度。
最大深度是指从根节点到最远叶子节点的最长路径上的节点总数。
N 叉树输入按层序遍历序列化表示，每组子节点由空值分隔（请参见示例）。
 */
type Node struct {
	Val      int
	Children []*Node
}
func maxDepth(root *Node) int {
	ans := 0
	if root == nil {
		return 0
	}
	queue := []*Node{root}
	for len(queue) > 0 {
		q := queue
		queue = nil
		for _, node := range q {
			queue = append(queue,node.Children...)
		}
		ans++
	}
	return ans
}
func maxDepth1(root *Node) int {
	if root == nil {
		return 0
	}
	ans := 0
	for _, child := range root.Children {
		if childDepth := maxDepth(child); childDepth > ans {
			ans = childDepth
		}
	}
	return ans+1
}
func maxDepth3(root *Node) int {
	if root == nil {
		return 0
	}
	max := 0
	for _, v := range root.Children {
		dp := maxDepth3(v)
		if max < dp{
			max = dp
		}
	}
	return max+1
}