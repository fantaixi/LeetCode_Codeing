package main

import "container/list"

func main() {

}

type Node struct {
	Val      int
	Children []*Node
}
/*
给定一个 N 叉树，返回其节点值的层序遍历。（即从左到右，逐层遍历）。
树的序列化输入是用层序遍历，每组子节点都由 null 值分隔（参见示例）。
 */
func levelOrder(root *Node) [][]int {
	res := [][]int{}
	queue := list.New()
	if root == nil {
		return res
	}
	queue.PushBack(root)
	for queue.Len() >0 {
		length := queue.Len()
		temArr := []int{} //
		for j := 0; j < length; j++ { //该层的每个元素：一添加到该层的结果集中；二找到该元素的下层元素加入到队列中，方便下次使用
			myNode := queue.Remove(queue.Front()).(*Node)
			temArr = append(temArr,myNode.Val)
			for i := 0; i < len(myNode.Children); i++ {
				queue.PushBack(myNode.Children[i])
			}
		}
		res = append(res,temArr)
	}
	return res
}
