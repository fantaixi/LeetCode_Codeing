package main

func main() {

}
func maxProfit(prices []int, fee int) int {
	result,minPrice := 0,prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
		// 保持状态，不买不卖，因为卖出后当次收益小于0
		if prices[i] >= minPrice && prices[i] <= minPrice + fee{
			continue
		}
		if prices[i] > minPrice + fee {
			result += prices[i] - minPrice - fee
			// 更新低价，如遇到更高价格，抵消fee，
			// 相当于阶段性的最低买最高卖, 可实现持续累加，最低买最高卖，仅一次手续费
			minPrice = prices[i] - fee
		}
	}
	return result
}