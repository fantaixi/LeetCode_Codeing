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
给你一棵二叉树的根节点 root ，返回其节点值的 后序遍历 。
 */
/*
递归
 */
func postorderTraversal(root *TreeNode) []int {
	res := []int{}
	var tra func(node *TreeNode)
	tra = func(node *TreeNode) {
		if node == nil {
			return
		}
		tra(node.Left)
		tra(node.Right)
		res = append(res,node.Val)
	}
	tra(root)
	return res
}

/*
遍历
先序遍历是中左右，后续遍历是左右中，那么我们只需要调整一下先序遍历的代码顺序，
就变成中右左的遍历顺序，然后在反转result数组，输出的结果顺序就是左右中
 */
func postorderTraversal1(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return nil
	}
	st := list.New()
	st.PushBack(root)
	for st.Len() > 0{
		node := st.Remove(st.Back()).(*TreeNode)
		res = append(res,node.Val)
		if node.Left != nil {
			st.PushBack(node.Left)
		}
		if node.Right != nil {
			st.PushBack(node.Right)
		}
	}
	reverse(res)
	return res
}

func reverse(a []int) {
	left,right := 0,len(a)-1
	for left < right {
		a[left],a[right] = a[right],a[left]
		left++
		right--
	}
}

/*
统一
 */
func postorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := []int{}
	stack := list.New()
	stack.PushBack(root)
	var node *TreeNode
	for stack.Len() > 0 {
		e := stack.Back()
		stack.Remove(e)
		if e.Value == nil {
			e = stack.Back()
			stack.Remove(e)
			node = e.Value.(*TreeNode)
			res = append(res,node.Val)
			continue
		}
		node = e.Value.(*TreeNode)
		stack.PushBack(node)
		stack.PushBack(nil)
		if node.Right != nil {
			stack.PushBack(node.Right)
		}
		if node.Left != nil {
			stack.PushBack(node.Left)
		}
	}
	return res
}