package main

import "fmt"

func str(s string) (string, int) {
	memo := make(map[string]int)
	for _, v := range s {
		memo[string(v)]++
	}
	max := 0
	str := ""
	for char, count := range memo {
		if count > max {
			max = count
			str = char
		}
	}
	return str, max
}
func main() {
	fmt.Println(str("meles"))

}