package main

import "fmt"

func romanToInt(s string) int {
	rTi := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	total := 0
	for i := 0; i < len(s); i++ {
		if i < len(s)-1 && rTi[s[i]] < rTi[s[i+1]] {
			total -= rTi[s[i]]
		} else {
			total += rTi[s[i]]
		}
	}
	return total
}
func main() {
	fmt.Println(romanToInt("XXXX"))
}