package _01

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
给定一个有相同值的二叉搜索树（BST），找出 BST 中的所有众数（出现频率最高的元素）。
假定 BST 有如下定义：
结点左子树中所含结点的值小于等于当前结点的值
结点右子树中所含结点的值大于等于当前结点的值
左子树和右子树都是二叉搜索树
 */
/*
可以顺序扫描中序遍历序列，用 base 记录当前的数字，用 count 记录当前数字重复的次数，
用 maxCount 来维护已经扫描过的数当中出现最多的那个数字的出现次数，用 res 数组记录出现的众数
首先更新 base 和 count:
如果该元素和 base 相等，那么 count 自增 1；
否则将 base 更新为当前数字，count 复位为 1。
然后更新 maxCount：
如果 count=maxCount，那么说明当前的这个数字（base）出现的次数等于当前众数出现的次数，将 base 加入 answer 数组；
如果 count>maxCount，那么说明当前的这个数字（base）出现的次数大于当前众数出现的次数，
因此，我们需要将 maxCount 更新为 count，清空 res 数组后将base 加入 answer 数组。
 */
func findMode(root *TreeNode) []int {
	res := []int{}
	var base,count,maxCount int
	update := func(x int) {
		if x == base {
			count++
		}else {
			base,count = x,1
		}
		if count == maxCount {
			res = append(res,base)
		}else if count > maxCount {
			maxCount = count
			res = []int{base}
		}
	}
	dfs := func(node *TreeNode){}
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		update(node.Val)
		dfs(node.Right)
	}
	dfs(root)
	return res
}
/*
Morris 中序遍历

记作当前节点为cur。
如果cur无左孩子，cur向右移动（cur=cur.right）
如果cur有左孩子，找到cur左子树上最右的节点，记为mostright
如果mostright的right指针指向空，让其指向cur，cur向左移动（cur=cur.left）
如果mostright的right指针指向cur，让其指向空，cur向右移动（cur=cur.right）

 */
func findMode1(root *TreeNode) []int {
	res := []int{}
	var base,count,maxCount int
	update := func(x int) {
		if x == base {
			count++
		}else {
			base,count = x,1
		}
		if count == maxCount {
			res = append(res,base)
		}else if count > maxCount {
			maxCount = count
			res = []int{base}
		}
	}
	cur := root
	for cur != nil {
		//如果cur无左孩子，cur向右移动（cur=cur.right）
		if cur.Left == nil {
			update(cur.Val)
			cur = cur.Right
			continue
		}
		pre := cur.Left
		//找到左子树上最右的结点，并让其指向cur
		for pre.Right != nil && pre.Right != cur {
			pre = pre.Right
		}
		//如果mostright的right指针指向空，让其指向cur，cur向左移动（cur=cur.left）
		if pre.Right == nil {
			pre.Right = cur
			cur = cur.Left
		}else {  //如果mostright的right指针指向cur，让其指向空，cur向右移动（cur=cur.right）
			pre.Right = cur
			update(cur.Val)
			cur = cur.Right
		}
	}
	return res
}