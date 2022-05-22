package main

import (
	"container/list"
	"math"
)

func main() {

}
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
/*
给定一棵二叉树的根节点 root ，请找出该二叉树中每一层的最大值。
 */
func largestValues(root *TreeNode) []int {
	res := [][]int{}
	finalArr := []int{}
	queue := list.New()
	if root == nil {
		return finalArr
	}
	queue.PushBack(root)
	temp := []int{}
	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			myNode := queue.Remove(queue.Front()).(*TreeNode)
			if myNode.Left != nil {
				queue.PushBack(myNode.Left)
			}
			if myNode.Right != nil {
				queue.PushBack(myNode.Right)
			}
			temp = append(temp,myNode.Val)
		}
		res = append(res,temp)
		temp = []int{}
	}
	finalLen := len(res)
	for i := 0; i < finalLen; i++ {
		finalArr = append(finalArr,max(res[i]...))
	}
	return finalArr
}
func max(val ...int) int {
	max := int(math.Inf(-1))  //负无穷
	for _, v := range val {
		if v > max {
			max=v
		}
	}
	return max
}