package main

import "sort"

func main() {

}
func reconstructQueue(people [][]int) [][]int {
	//将身高从大到小排序
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			//当身高相同时，将K按照从小到大排序
			return people[i][1] < people[j][1]
		}
		//确保身高按照由大到小的顺序来排，并不确定K是按照从小到大排序的
		return people[i][0] > people[j][0]
	})
	res := make([][]int,0)
	for _, val := range people {
		res = append(res,val)
		//后移一位，腾出空间
		copy(res[val[1]+1:],res[val[1]:])
		res[val[1]] = val
	}
	return res
}
