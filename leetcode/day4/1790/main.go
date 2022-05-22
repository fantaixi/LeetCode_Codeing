package main

import "fmt"

func main() {
 	fmt.Println(areAlmostEqual("bank","kanb"))
 	fmt.Println(len("bank"))
}
func areAlmostEqual(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}
	count,i,j := 0,0,0
	flag := 0
	for i < len(s1){
		if s1[i] != s2[i] {
			count++
			if count == 1 {
				j = i
			}
			if count == 2 {
				if s1[i] == s2[j] && s1[j] == s2[i] {
					flag = 1
				}
			}
			if count > 2 {
				return false
			}
		}
		i++
	}
	if flag==1 {
		return true
	}
	return false
}