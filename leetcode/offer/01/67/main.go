package main

import "math"

func main() {

}
func strToInt(str string) int {
	sum,sign,i,n := 0,1,0,len(str)
	//去掉空格
	for i < n && str[i] == ' ' {
		i++
	}
	//判断正负
	if i<n {
		if str[i] == '-' {
			sign = -1
			i++
		}else if str[i] == '+' {
			sign = 1
			i++
		}
	}
	//从左到右依次累加
	for i < n && str[i] >= '0' && str[i] <= '9' {
		sum = 10*sum+int(str[i]-'0')
		if sign*sum < math.MinInt32 {
			return math.MinInt32
		}else if sign*sum > math.MaxInt32 {
			return math.MaxInt32
		}
		i++
	}
	return sum*sign
}