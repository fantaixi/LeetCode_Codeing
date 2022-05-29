package _131

/*
上一版因为每次都会调用isPartition，会造成空间的浪费，所以每次都将结果存起来，每次遇到一样的就不需要再调用isPartition
 */
func partition1(s string) [][]string {
	var path []string
	var res [][]string
	// 0表示未计算，1表示是回文，2表示不是回文
	memo := make([][]int,len(s))
	for i := range memo {
		memo[i] = make([]int,len(s))
	}
	backtracking(s,path,0,&res,memo)
	return res
}
func backtracking(s string, path []string, startIndex int, res *[][]string,memo [][]int) {
	//到达字符串末尾了
	if startIndex == len(s) {
		t := make([]string,len(path))
		copy(t,path)
		*res = append(*res,t)
	}

	for i := startIndex; i < len(s); i++ {
		//不是回文，直接跳过
		if memo[startIndex][i] == 2 {
			continue
		}
		//记录为回文 或没有记录但isPali调用为真
		if memo[startIndex][i] == 1 || isParl(s, startIndex, i, memo) {
			path = append(path,s[startIndex:i+1])
			backtracking(s,path,i+1,res,memo)
			path = path[:len(path)-1]
		}
	}
}
//判断是否为回文
func isParl(s string,left,right int,memo [][]int) bool {
	//因为left和right都不是固定的，所以不能直接赋值为0和len(s)
	for left < right {
		if s[left] != s[right] {
			memo[left][right] = 2  //存入memo
			return false
		}
		left++
		right--
	}
	memo[left][right] = 1  //存入memo
	return true
}
