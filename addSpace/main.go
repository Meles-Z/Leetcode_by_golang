package main

import "fmt"

func addSpaces(s string, spaces []int) string {
	result := ""
	spaceIndex := 0
	for i := 0; i < len(s); i++ {
		if spaceIndex < len(spaces) && i == spaces[spaceIndex] {
			result += " "
			spaceIndex++
		}
		result+=string(s[i])
	}
	return result
}

func main() {
	s := "LeetcodeHelpsMeLearn"
	spaces := []int{8, 13, 15}
	fmt.Println(addSpaces(s, spaces))
}
