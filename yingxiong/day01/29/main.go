package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(-math.Pow(2,31))
	fmt.Println(divide1(-2147483648,-1))
}
func divide(dividend int, divisor int) int {
	var abs func(int)int
	abs = func(a int) int{
		if a<0{
			return -a
		}
		return a
	}
	//是否为负数
	sign:= (dividend >0&&divisor<0)||(dividend<0 && divisor>0)
	dividend = abs(dividend)
	divisor = abs(divisor)
	ans:=0
	for dividend>=divisor{
		//d 除数 q 商
		d,q:=divisor,1
		for d<=dividend{
			d,q=d<<1,q<<1
		}
		ans+=q>>1
		dividend = dividend-(d>>1)
	}
	if sign{
		ans = ^ans+1
	}

	if ans < -2147483648||ans>2147483647{
		return 2147483647
	}
	return ans
}
func divide1(dividend int, divisor int) int {
	if(dividend== -2147483648 && divisor == -1) {
		return 2147483647;
	}
	return dividend / divisor
}