package main

import "container/list"

func main() {

}
//与117等同
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next *Node
}
/*
给定一个完美二叉树，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下：
struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。
初始状态下，所有next 指针都被设置为 NULL。
 */
func connect(root *Node) *Node {
	res := [][]*Node{}
	if root == nil {
		return root
	}
	queue := list.New()
	queue.PushBack(root)
	var temArr []*Node
	for queue.Len() > 0 {
		length := queue.Len() //保存当前层的长度，然后处理当前层（十分重要，防止添加下层元素影响判断层中元素的个数）
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*Node)  //出队列
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			temArr =append(temArr,node)  //将值加入本层切片中
		}
		res=append(res, temArr) //放入结果集
		temArr = []*Node{}  //清空层的数据
	}
	//指定next
	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res[i])-1; j++ {
			res[i][j].Next = res[i][j+1]
		}
	}
	return root
}