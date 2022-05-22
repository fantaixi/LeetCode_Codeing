package _42

/*
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
 */
/*
hash: 将 s[i] - ‘a’ 所在的元素做v+1，再遍历t的时候，如果有相同的字符就做v-1，结束之后，如果有元素不为0，返回false，全为0则返回true
 */
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	res := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if v,ok := res[s[i]];v >=0 && ok {
			res[s[i]] = v + 1
		}else {
			res[s[i]] = 1
		}
	}
	for i := 0; i < len(t); i++ {
		if v, ok := res[t[i]]; v >= 1 && ok {
			res[t[i]] = v-1
		}else {
			return false
		}
	}
	return true
}

func isAnagram1(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	cnt := map[rune]int{}
	for _, v := range s {
		cnt[v]++
	}
	for _, v := range t {
		cnt[v]--
		if cnt[v] < 0 {
			return false
		}
	}
	return true
}

func isAnagram2(s string, t string) bool {
	var c1,c2 [26]int
	for _, v := range s {
		c1[v-'a']++
	}
	for _, v := range t {
		c2[v-'a']++
	}
	return c1 == c2
}
