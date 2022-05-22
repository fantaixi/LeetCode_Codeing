package main

import (
	"math"
	"sort"
	"strconv"
)

func main() {

}
func findMinDifference(timePoints []string) int {
	if len(timePoints) > 1440 {
		return 0
	}
	cache,min := make([]int,len(timePoints)),math.MaxInt32
	for i := 0; i < len(timePoints); i++ {
		hour,_ := strconv.Atoi(timePoints[i][0:2])
		minute,_ := strconv.Atoi(timePoints[i][3:])
		cache[i] = hour*60 + minute
	}
	sort.Ints(cache)
	for i := 1; i < len(cache); i++ {
		min = int(math.Min(float64(min),float64(cache[i] - cache[i-1])))
	}
	return int(math.Min(float64(min),float64(cache[0]+1440 - cache[len(cache)-1])))
}