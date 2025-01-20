package main

import "fmt"

func reverseString(str string) string {
	strr := ""

	for i := len(str) - 1; i >= 0; i-- {
		by := []rune(str)
		rev := by[i]
		strr += string(rev)
	}
	return strr
}
func main() {
	fmt.Println(reverseString("strops"))
}
