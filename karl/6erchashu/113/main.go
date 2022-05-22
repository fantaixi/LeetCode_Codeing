package _13
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
/*
给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。
叶子节点 是指没有子节点的节点。
 */
func pathSum(root *TreeNode, targetSum int) [][]int {
	res := [][]int{}
	dfs(root,&res,new([]int),targetSum)
	return res
}
func dfs(node *TreeNode, result *[][]int, currPath *[]int, targetSum int) {
	if node == nil {
		return
	}
	// 将targetSum在遍历每层的时候都减去本层节点的值
	targetSum -= node.Val
	// 把当前节点放到路径记录里
	*currPath = append(*currPath,node.Val)
	// 如果剩余的targetSum为0, 则正好就是符合的结果
	if node.Left == nil && node.Right == nil && targetSum == 0 {
		// 不能直接将currPath放到result里面, 因为currPath是共享的, 每次遍历子树时都会被修改
		temp := make([]int,len(*currPath))
		for i, val := range *currPath {
			temp[i] = val
		}
		// 将副本放到结果集里
		*result = append(*result,temp)
	}
	dfs(node.Left,result,currPath,targetSum)
	dfs(node.Right,result,currPath,targetSum)
	// 当前节点遍历完成, 从路径记录里删除掉
	*currPath = (*currPath)[:len(*currPath)-1]
}