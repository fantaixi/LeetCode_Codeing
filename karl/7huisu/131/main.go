package _131

/*
给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。
回文串 是正着读和反着读都一样的字符串。
 */

func partition(s string) [][]string {
	var path []string
	var res [][]string
	backtrack(s,path,0,&res)
	return res
}
func backtrack(s string, path []string, startIndex int, res *[][]string) {
	//到达字符串末尾了
	if startIndex == len(s) {
		t := make([]string,len(path))
		copy(t,path)
		*res = append(*res,t)
	}

	for i := startIndex; i < len(s); i++ {
		//处理（首先通过startIndex和i判断切割的区间，进而判断该区间的字符串是否为回文，
		//若为回文，则加入到path，否则继续后移，找到回文区间）（这里为一层处理）
		if isPartition(s, startIndex, i) {
			path = append(path,s[startIndex:i+1])
		}else {
			continue
		}
		backtrack(s,path,i+1,res)
		path = path[:len(path)-1]
	}
}
//判断是否为回文
func isPartition(s string,left,right int) bool {
	//因为left和right都不是固定的，所以不能直接赋值为0和len(s)
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}