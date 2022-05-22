package _83

/*
给你两个字符串：ransomNote 和 magazine ，判断 ransomNote 能不能由 magazine 里面的字符构成。
如果可以，返回 true ；否则返回 false 。
magazine 中的每个字符只能在 ransomNote 中使用一次。
 */
func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}
	cnt := [26]int{}
	for _, v := range magazine {
		cnt[v - 'a']++
	}
	for _, v := range ransomNote {
		cnt[v - 'a']--
		if cnt[v-'a'] < 0 {
			return false
		}
	}
	return true
}
