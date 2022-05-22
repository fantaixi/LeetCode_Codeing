package main

import "container/list"

func main() {

}
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
/*
给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
 */
func rightSideView(root *TreeNode) []int {
	res := [][]int{}
	queue := list.New()
	finalRes := []int{}
	if root == nil {
		return finalRes
	}
	queue.PushBack(root)
	for queue.Len() >  0{
		length := queue.Len()
		tem := []int{}
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			tem = append(tem,node.Val)
		}
		res = append(res,tem)
	}
	//取每一层的最后一个元素
	for i := 0; i < len(res); i++ {
		finalRes=append(finalRes,res[i][len(res[i])-1])
	}
	return finalRes
}