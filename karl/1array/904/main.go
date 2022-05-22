package main

func main() {

}
func totalFruit(fruits []int) int {
	//窗口类水果种类，及数量
	window := map[int]int{}
	length := len(fruits)
	start,end := 0,0
	count := 0
	for end < length {
		//窗口扩大
		window[fruits[end]]++
		end++
		//如果窗口内水果种类大于2，则缩小窗口
		for len(window) > 2 {
			window[fruits[start]]--
			if window[fruits[start]] == 0 {
				delete(window,fruits[start])
			}
			start++
		}
		sum := 0
		for _, i := range window {
			sum += i
		}
		count = max(count,sum)
	}
	return count
}
func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}