package main

import "fmt"

func reverseString(str string) string {
	strr := ""
	for i := len(str) - 1; i >= 0; i-- {
		rev := str[i]
		strr += string(rev)
	}
	return strr
}
func main() {
	fmt.Println(reverseString("Meles"))
}
