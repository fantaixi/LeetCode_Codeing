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
给你二叉树的根节点 root ，返回其节点值的 层序遍历 。
（即逐层地，从左到右访问所有节点）。
*/
func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	//返回一个初始化的列表。
	queue := list.New()
	// PushBack在列表l的后面插入一个新元素e，值为v，并返回e。
	queue.PushBack(root)
	var temArr []int
	for queue.Len()>0 {
		length:=queue.Len()//保存当前层的长度，然后处理当前层（十分重要，防止添加下层元素影响判断层中元素的个数）
		for i:=0;i<length;i++{
			// 如果e是列表l的一个元素，Remove将e从l中移除。它返回元素的值e.Value。该元素不能是nil。
			//Front返回列表l的第一个元素，如果列表为空则返回nil。
			node:=queue.Remove(queue.Front()).(*TreeNode)//出队列
			if node.Left!=nil{  //左
				queue.PushBack(node.Left)
			}
			if node.Right!=nil{  //右
				queue.PushBack(node.Right)
			}
			temArr=append(temArr,node.Val)// 中  将值加入本层切片中
		}
		res=append(res,temArr)//放入结果集
		temArr=[]int{}//清空层的数据
	}
	return res
}