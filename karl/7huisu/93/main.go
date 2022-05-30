package _3

import "strconv"

/*
给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。
有效的 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
例如："0.1.2.201" 和 "192.168.1.1" 是 有效的 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效的 IP 地址。
*/

func restoreIpAddresses(s string) []string {
	var res,path []string
	backtrack(s,path,0,&res)
	return res
}
func backtrack(s string, path []string, startIndex int, res *[]string) {
	if startIndex == len(s) && len(path) == 4 {
		temp := path[0]+"."+path[1]+"."+path[2]+"."+path[3]
		*res = append(*res,temp)
	}
	for i := startIndex; i < len(s); i++ {
		path := append(path,s[startIndex:i+1])
		if i-startIndex+1 <= 3 && len(path) <= 4 && isValid(s, startIndex, i) {
			backtrack(s,path,i+1,res)
		}else {
			//如果首尾超过了3个，或路径多余4个，或前导为0，或大于255，直接回退
			return
		}
		path = path[:len(path)-1]
	}
}
/*
段位以0为开头的数字不合法
段位里有非正整数字符不合法
段位如果大于255了不合法
 */
func isValid(s string, start, end int) bool {
	checkInt, _ := strconv.Atoi(s[start : end+1])
	//对于前导 0的IP（特别注意s[startIndex]=='0'的判断，不应该写成s[startIndex]==0，因为s截取出来不是数字）
	if end-start+1 > 1 && s[start] == '0' {
		return false
	}
	if checkInt > 255 {
		return false
	}
	return true
}
