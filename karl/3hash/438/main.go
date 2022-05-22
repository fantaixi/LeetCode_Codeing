package _38

/*
给定两个字符串s和 p，找到s中所有p的异位词的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
异位词 指由相同字母重排列形成的字符串（包括相同的字符串）。
*/
/*
滑动窗口
每次比较26位的区别是没有必要的，因为变动总是一个字母，一位的变动，
可以维护一个当前窗口和p的统计的不一样的个数，维护在同一个哈希表里，
当某位变成零，不一样的个数减一；当某位从零变成非零，不一样的个数加一。
只有当不一样的个数为0时，才是想要的答案
*/
func findAnagrams(s string, p string) []int {
	ans := []int{}
	m, n := len(s), len(p)
	if m < n {
		return ans
	}
	var cntS, cntP [26]int
	for i := 0; i < n; i++ {
		cntP[p[i]-byte('a')] ++
	}
	for i := 0; i < m; i++ {
		cntS[s[i]-byte('a')]++
		if i >= n-1 {  // 一次比较一位
			if cntS == cntP {
				ans = append(ans, i-n+1)
			}
			cntS[s[i-n+1]-byte('a')]--
		}
	}
	return ans
}

func findAnagrams1(s string, p string) []int {
	ans := []int{}
	sLen, pLen := len(s), len(p)
	if sLen < pLen {
		return ans
	}
	var sCnt, pCnt [26]int
	for i, v := range p {
		sCnt[s[i]-'a']++
		pCnt[v-'a']++
	}
	if sCnt == pCnt {
		ans = append(ans,0)
	}
	for i, v := range s[:sLen-pLen] {
		sCnt[v-'a']--
		sCnt[s[i+pLen]-'a']++
		if sCnt == pCnt {
			ans = append(ans,i+1)
		}
	}
	return ans
}
