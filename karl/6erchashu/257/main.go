package _57

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
/*
给你一个二叉树的根节点 root ，按 任意顺序 ，返回所有从根节点到叶子节点的路径。
叶子节点 是指没有子节点的节点。
*/
func binaryTreePaths(root *TreeNode) []string {
	res := []string{}
	build(root,"",&res)
	return res
}
func build(root *TreeNode, pathStr string, res *[]string) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		pathStr += strconv.Itoa(root.Val)
		*res = append(*res,pathStr)
		return
	}
	pathStr += strconv.Itoa(root.Val) + "->"
	build(root.Left,pathStr,res)
	build(root.Right,pathStr,res)
}

func binaryTreePaths1(root *TreeNode) []string {
	res := []string{}
	var dfs func(node *TreeNode,s string)
	dfs = func(node *TreeNode, s string) {
		if node.Left == nil && node.Right == nil {
			v := s + strconv.Itoa(node.Val)
			res = append(res,v)
			return
		}
		s += strconv.Itoa(node.Val) + "->"
		if node.Left != nil {
			dfs(node.Left,s)
		}
		if node.Right != nil {
			dfs(node.Right,s)
		}
	}
	dfs(root,"")
	return res
}