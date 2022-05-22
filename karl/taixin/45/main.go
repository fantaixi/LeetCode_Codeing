package main

func main() {

}
func jump1(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	//当前覆盖最远距离下标，记录走的最大步数，下一步覆盖最远距离下标
	curDistance,ans,nextDistance := 0,0,0
	for i := 0; i < len(nums); i++ {
		// 更新下⼀步覆盖最远距离下标
		nextDistance = max(nums[i]+i,nextDistance)
		// 遇到当前覆盖最远距离下标
		if i == curDistance {
			// 如果当前覆盖最远距离下标不是终点
			if curDistance != len(nums)-1 {
				ans++
				// 更新当前覆盖最远距离下标
				curDistance = nextDistance
				// 下⼀步的覆盖范围已经可以达到终点，结束循环
				if nextDistance >= len(nums)-1 {
					break
				}
			}else {
				break
			}
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func jump(nums []int) int {
	//当前覆盖最远距离下标，记录走的最大步数，下一步覆盖最远距离下标
	curDistance,ans,nextDistance := 0,0,0
	for i := 0; i < len(nums)-1; i++ {
		nextDistance = max(nums[i]+i,nextDistance)
		if i == curDistance {
			curDistance = nextDistance
			ans++
		}
	}
	return ans
}