package main

func main() {

}
/*
准备两个指针 i, j 分别指向 S，T 的末位字符，再准备两个变量 skipS，skipT 来分别存放 S，T 字符串中的 # 数量。
从后往前遍历 S，所遇情况有三，如下所示：
2.1 若当前字符是 #，则 skipS 自增 1；
2.2 若当前字符不是 #，且 skipS 不为 0，则 skipS 自减 1；
2.3 若当前字符不是 #，且 skipS 为 0，则代表当前字符不会被消除，我们可以用来和 T 中的当前字符作比较。
若对比过程出现 S, T 当前字符不匹配，则遍历结束，返回 false，若 S，T 都遍历结束，且都能一一匹配，则返回 true。
*/
func backspaceCompare(s string, t string) bool {
	spaceS,spaceT := 0,0
	i,j := len(s)-1,len(t)-1
	for i >= 0 || j >= 0{
		for i >= 0 {
			if s[i] == '#' {
				spaceS++
				i--
			}else if spaceS > 0 {
				spaceS--
				i--
			}else {
				break
			}
		}
		for j >= 0 {
			if t[j] == '#' {
				spaceT++
				j--
			}else if spaceT > 0 {
				spaceT--
				j--
			}else {
				break
			}
		}
		if i >= 0 && j >= 0 {
			if s[i] != t[j] {
				return false
			}
		}else if i >= 0 || j >= 0 {
			return false
		}
		i--
		j--
	}
	return true
}