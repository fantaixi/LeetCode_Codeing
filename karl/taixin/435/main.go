package main

import "sort"

func main() {

}
func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	//先根据右边界排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	n,ans,right := len(intervals),1,intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		//如果起始点 大于等于right时，不重叠
		if intervals[i][0] >= right {
			ans++
			//更新right的位置
			right = intervals[i][1]
		}
	}
	//需要移除的数量= 总数量-不重叠的数量
	return n - ans
}