package main

import (
	"regexp"
	"strings"
)

func main() {

}
func isNumber(s string) bool {
	s = strings.TrimSpace(s)
	res,_ := regexp.MatchString("^(([\\+\\-]?[0-9]+(\\.[0-9]*)?)|([\\+\\-]?\\.?[0-9]+))((e|E)[\\+\\-]?[0-9]+)?$", s)
	return res
}