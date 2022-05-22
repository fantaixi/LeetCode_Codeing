package main

func main() {

}
func replaceSpace(s string) string {
	newStr := ""
	for _, val := range s {
		if val == ' ' {
			newStr += "%20"
		}else {
			newStr += string(val)
		}
	}
	return newStr
}