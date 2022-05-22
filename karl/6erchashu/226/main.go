package main

import "container/list"

func main() {
}
/*
给你一棵二叉树的根节点 root ，翻转这棵二叉树，并返回其根节点。
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
/*
递归的前序遍历
 */
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	temp := root.Left
	root.Left = root.Right
	root.Right = temp
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}
/*
递归的后序遍历
*/
func invertTree1(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	invertTree(root.Left) //遍历左节点
	invertTree(root.Right)  //遍历右节点
	root.Left,root.Right = root.Right,root.Left  //交换
	return root
}
/*
迭代的前序遍历
*/
func invertTree2(root *TreeNode) *TreeNode {
	stack := []*TreeNode{}
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			node.Left,node.Right = node.Right,node.Left  //交换
			stack = append(stack,node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		node=node.Right
	}
	return root
}

/*
迭代的后序遍历
*/
func invertTree3(root *TreeNode) *TreeNode {
	stack := []*TreeNode{}
	node := root
	var prev *TreeNode
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack=append(stack,node)
			node=node.Left
		}
		node=stack[len(stack)-1]
		stack=stack[:len(stack)-1]
		if node.Right == nil || node.Right == prev {
			node.Left,node.Right = node.Right,node.Left
			prev = node
			node= nil
		}else {
			stack= append(stack,node)
			node=node.Right
		}
	}
	return root
}

/*
层序遍历
 */
func invertTree4(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	queue := list.New()
	node := root
	queue.PushBack(node)
	for queue.Len()>0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			e := queue.Remove(queue.Front()).(*TreeNode)
			e.Left,e.Right = e.Right,e.Left
			if e.Left != nil {
				queue.PushBack(e.Left)
			}
			if e.Right != nil {
				queue.PushBack(e.Right)
			}
		}
	}
	return root
}