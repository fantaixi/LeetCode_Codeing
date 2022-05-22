package main

import "strings"

func main() {

}
func replaceSpace(s string) string {
	newArr := ""
	for _, v := range s {
		if v == ' ' {
			newArr += "%20"
		}else {
			newArr += string(v)
		}
	}
	return newArr
}

func replaceSpace1(s string) string {
	return strings.Replace(s," ","%20",-1)
}
