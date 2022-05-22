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
给定一个非空二叉树的根节点 root , 以数组的形式返回每一层节点的平均值。
与实际答案相差 10-5 以内的答案可以被接受。
 */
func averageOfLevels(root *TreeNode) []float64 {
	res := [][]int{}
	finalRes := []float64{}
	if root == nil {
		return finalRes
	}
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
			tempArr = append(tempArr,node.Val)
		}
		res = append(res,tempArr)
		tempArr=[]int{}
	}
	//计算平均值
	length := len(res)
	for i := 0; i < length; i++ {
		sum := 0  //这里sum不能放在外面，否则会丢失精度
		for j := 0; j < len(res[i]); j++ {
			sum += res[i][j]
		}
		temp := float64(sum)/float64(len(res[i]))
		finalRes = append(finalRes,temp)
	}
	return finalRes
}