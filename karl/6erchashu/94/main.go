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
给定一个二叉树的根节点 root ，返回 它的 中序 遍历 。
 */
func inorderTraversal(root *TreeNode) []int {
	res :=[]int{}
	var tra func(node *TreeNode)
	tra = func(node *TreeNode) {
		if node == nil {
			return
		}
		tra(node.Left)
		res = append(res,node.Val)
		tra(node.Right)
	}
	tra(root)
	return res
}
/*
遍历：

 */

func inorderTraversal1(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return nil
	}
	st := list.New()
	cur := root  //方便理解（换成root也可以）
	for cur != nil || st.Len() > 0 {
		if cur != nil {
			st.PushBack(cur)  // 将访问的节点放进栈
			cur = cur.Left  //左
		}else {
			cur = st.Remove(st.Back()).(*TreeNode) // 从栈里弹出的数据，就是要处理的数据
			res = append(res,cur.Val)  //中
			cur = cur.Right  //右
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
		//压栈顺序： 右中左
		if node.Right != nil {
			stack.PushBack(node.Right)
		}
		stack.PushBack(node) //中间节点压栈后再压入nil作为中间节点的标志符
		stack.PushBack(nil)
		if node.Left != nil {
			stack.PushBack(node.Left)
		}

	}
	return res
}