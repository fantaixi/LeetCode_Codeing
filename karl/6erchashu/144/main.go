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
给你二叉树的根节点 root ，返回它节点值的 前序 遍历。
 */
/*
递归
 */
func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	var tra func(node *TreeNode)
	// 1.递归函数及参数确定
	tra = func(node *TreeNode) {
		// 2.递归终止条件
		if node == nil {
			return
		}
		//处理父节点
		res = append(res,node.Val)
		//处理左子树
		tra(node.Left)
		//处理右子树
		tra(node.Right)
	}
	tra(root)
	return res
}

/*
遍历:
前序遍历是中左右，
每次先处理的是中间节点，
那么先将根节点放入栈中，然后将右孩子加入栈，再加入左孩子。
 */
func preorderTraversal1(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return nil
	}
	st := list.New()
	//PushBack在列表l的后面插入一个新元素e，值为v，并返回e。
	st.PushBack(root)
	for st.Len() > 0 {
		// Remove:如果e是列表l的一个元素，Remove将e从l中移除。它返回元素的值e.Value。该元素不能是nil。
		// Back:返回列表l的最后一个元素，如果列表为空，则返回nil。
		node := st.Remove(st.Back()).(*TreeNode)
		res = append(res,node.Val)
		if node.Right != nil {
			st.PushBack(node.Right)
		}
		if node.Left != nil {
			st.PushBack(node.Left)
		}
	}
	return res
}

/*
统一遍历
type Element struct {
    // 元素保管的值
    Value interface{}
    // 内含隐藏或非导出字段
}

func (l *List) Back() *Element
前序遍历：中左右
压栈顺序：右左中
 */
func preorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := []int{}
	stack := list.New()  //栈
	stack.PushBack(root)
	var node *TreeNode
	for stack.Len() > 0 {
		e := stack.Back()
		stack.Remove(e) //弹出元素
		if e.Value == nil { // 如果为空，则表明是需要处理中间节点
			e = stack.Back() //弹出元素（即中间节点）
			stack.Remove(e)  //删除中间节点
			node = e.Value.(*TreeNode)
			res = append(res,node.Val)  //将中间节点加入到结果集中
			continue  //继续弹出栈中下一个节点
		}
		node = e.Value.(*TreeNode)
		//压栈顺序：右左中
		if node.Right != nil {
			stack.PushBack(node.Right)
		}
		if node.Left != nil {
			stack.PushBack(node.Left)
		}
		stack.PushBack(node) //中间节点压栈后再压入nil作为中间节点的标志符
		stack.PushBack(nil)
	}
	return res
}