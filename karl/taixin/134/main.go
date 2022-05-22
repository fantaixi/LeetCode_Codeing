package main

func main() {

}
func canCompleteCircuit(gas []int, cost []int) int {
	start,totalSum,curSum := 0,0,0
	for i := 0; i < len(gas); i++ {
		curSum += gas[i]-cost[i]
		totalSum += gas[i]-cost[i]
		// 当前累加rest[i]和 curSum⼀旦⼩于0
		if curSum < 0 {
			// 起始位置更新为i+1
			start = i +1
			// curSum从0开始
			curSum = 0
		}
	}
	// 说明怎么⾛都不可能跑⼀圈
	if totalSum < 0 {
		return -1
	}
	return start
}