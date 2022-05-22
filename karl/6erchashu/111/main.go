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
给定一个二叉树，找出其最小深度。
最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
说明：叶子节点是指没有子节点的节点。

只有当左右孩子都为空的时候，才说明遍历的最低点了。如果其中一个孩子为空则不是最低点
 */
/*
如果左子树为空，右子树不为空，说明最小深度是 1 + 右子树的深度。
反之，右子树为空，左子树不为空，最小深度是 1 + 左子树的深度。
最后如果左右子树都不为空，返回左右子树深度最小值 + 1 。
 */
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right != nil {
		return minDepth(root.Right)+1
	}
	if root.Left != nil && root.Right == nil {
		return minDepth(root.Left)+1
	}
	return min(minDepth(root.Left),minDepth(root.Right))+1
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minDepth1(root *TreeNode) int {
	dep := 0
	queue := make([]*TreeNode,0)
	if root != nil {
		queue = append(queue,root)
	}
	for l := len(queue); l > 0; {
		dep++
		for ; l > 0; l-- {
			node := queue[0]
			if node.Left == nil && node.Right == nil {
				return dep
			}
			if node.Left != nil {
				queue = append(queue,node.Left)
			}
			if node.Right != nil {
				queue = append(queue,node.Right)
			}
			queue = queue[1:]
		}
		l = len(queue)
	}
	return dep
}

func minDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := 0
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		length :=queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left == nil && node.Right == nil {  //当前节点没有左右节点，则代表此层是最小层
				return ans+1  //返回当前层 ans代表是上一层
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		ans++  //记录层数
	}
	return ans+1
}