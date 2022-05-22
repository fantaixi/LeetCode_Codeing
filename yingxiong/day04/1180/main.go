package main

func main() {

}
func defangIPaddr(address string) string {
	newAdd := ""
	for _, val := range address {
		if val == '.' {
			newAdd += "[.]"
		}else {
			newAdd += string(val)
		}
	}
	return newAdd
}