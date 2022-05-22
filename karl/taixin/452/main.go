package main

import "sort"

func main() {

}
func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	res := 1
	for i := 1; i < len(points); i++ {
		//如果前一位的左边界小于后一位的右边界，则不重合
		if points[i-1][1] < points[i][0] {
			res++
		}else {
			//更新重叠气球最小右边界,覆盖该位置的值，留到下一步使用
			points[i][1] = min(points[i-1][1],points[i][1])
		}
	}
	return res
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}