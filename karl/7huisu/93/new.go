package _3

import (
	"strconv"
	"strings"
)

/*
换个写法
*/
func restoreIpAddresses1(s string) []string {
	res := []string{}
	var dfs func(tempStr []string, start int)
	dfs = func(tempStr []string, start int) {
		// 片段满4段，且耗尽所有字符,拼成字符串，加入解集
		if len(tempStr) == 4 && start == len(s) {
			// Join将其第一个参数中的元素连接起来，创建一个单一的字符串。字符串sep被放在结果字符串的元素之间。
			res = append(res, strings.Join(tempStr, "."))
			return
		}
		// 满4段，字符未耗尽，不用往下选了
		if len(tempStr) == 4 && start < len(s) {
			return
		}
		// 枚举出选择，三种切割长度
		for length := 1; length <= 3; length++ {
			// 加上要切的长度就越界，不能切这个长度
			if start+length-1 >= len(s) {
				return
			}
			// 不能切出'0x'、'0xx'
			if length != 1 && s[start] == '0' {
				return
			}
			// 当前选择切出的片段
			str := s[start : start+length]
			// 不能超过255
			if n, _ := strconv.Atoi(str); n > 255 {
				return
			}
			tempStr = append(tempStr, str)
			dfs(tempStr, start+length)
			tempStr = tempStr[:len(tempStr)-1]
		}
	}
	dfs([]string{}, 0)
	return res
}
