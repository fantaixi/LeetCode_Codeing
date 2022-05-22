package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Printf("第%d题：\n",i+1)
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
func numColor(root *TreeNode) int {
	m := make(map[int]struct{},0)
	dfs(root,m)
	return len(m)
}
func dfs(root *TreeNode, m map[int]struct{}) {
	if root == nil {
		return
	}
	m[root.Val]= struct{}{}
	dfs(root.Left,m)
	dfs(root.Right,m)
}