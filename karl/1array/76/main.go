package main

import "math"

func main() {

}
/*
给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。
如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
*/

/*
先遍历 t， 记录元素到哈希表 need 中；记录需要遍历的个数 len(t) 到 needCnt。
初始化长度为2的 ret 数组[0, Inf]，作为记录目标子字符串的左右索引下标。
遍历字符串 s， 如果遍历到的元素刚好>0, 说明是t中元素，needCnt减一。将遍历到的元素-1记录到哈希表 need 中。
如果 needCnt == 0 说明此时左右下标范围内 [lo,hi] 有符合要求子字符串, 开始缩小滑动窗口。
左索引下标 lo 向前推进, 当 need[s[lo]] == 0 时，说明遍历到t中元素（并且因为此时 needCntCnt==0）， s[lo,hi]为结果之一，判断是否最小。

*/
func minWindow(s string, t string) string {
	need := make(map[byte]int)
	for i := range t {
		need[t[i]]++
	}
	needCnt := len(t)
	ret := [2]int{0,math.MaxInt32}
	var lo int
	for hi := range s{
		if need[s[hi]] > 0 {
			needCnt--
		}
		need[s[hi]]--
		if needCnt == 0 {
			for true {
				if need[s[lo]] == 0 {
					break
				}
				need[s[lo]]++
				lo++
			}
			if hi-lo < ret[1]-ret[0] {
				ret = [2]int{lo,hi}
			}
			need[s[lo]]++
			needCnt++
			lo++
		}
	}
	if ret[1] > len(s) {
		return ""
	}
	return s[ret[0]:ret[1]+1]
}