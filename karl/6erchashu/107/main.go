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
给你二叉树的根节点 root ，返回其节点值 自底向上的层序遍历 。
（即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）
 */
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := [][]int{}
	queue := list.New()
	queue.PushBack(root)
	var tempArr []int
	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			tempArr=append(tempArr,node.Val)
		}
		res = append(res,tempArr)
		tempArr=[]int{}
	}
	reverse(res)
	return res
}
func reverse(arr [][]int) {
	left,right := 0,len(arr)-1
	for left < right{
		arr[left],arr[right] = arr[right],arr[left]
		left++
		right--
	}
}
