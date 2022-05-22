package _13

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
/*
给定一个二叉树的 根节点 root，请找出该二叉树的 最底层 最左边 节点的值。
假设二叉树中至少有一个节点。
 */
/*
迭代：返回最后一行第一个节点的数值
 */
func findBottomLeftValue(root *TreeNode) int {
	res := 0
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			//最后一行第一个节点
			if i == 0 {
				res = node.Val
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}
	return res
}

/*
递归
求深度，用前序遍历，求高度用后序遍历，这里用前序
 */
var maxDeep int //深度
var value int //最终值
func findBottomLeftValue1(root *TreeNode) int {
	//去除只有一个结点的情况
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	res := findLeftVal(root,maxDeep)
	return res
}
func findLeftVal(root *TreeNode, deep int) int{
	//最左边的值在左边
	if root.Left == nil && root.Right == nil {
		if deep > maxDeep {
			value = root.Val
			maxDeep = deep
		}
	}
	if root.Left != nil {
		deep++
		findLeftVal(root.Left,deep)
		deep--  //回溯
	}
	if root.Right != nil {
		deep++
		findLeftVal(root.Right,deep)
		deep-- //回溯
	}
	return value
}