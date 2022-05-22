package main

func main() {

}
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	pre := strs[0]
	for i := 1; i < len(strs); i++ {
		pre = lcp(pre,strs[i])
		if len(pre) == 0 {
			break
		}
	}
	return pre
}
func lcp(str1, str2 string) string {
	length := min(len(str1),len(str2))
	index := 0
	for index < length && str1[index] == str2[index] {
		index++
	}
	return str1[:index]
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}