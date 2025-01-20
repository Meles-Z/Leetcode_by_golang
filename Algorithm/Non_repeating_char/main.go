// Non repeating character - You are given a string consisting of only lower case and upper-case English Alphabets
//  and integers 0 to 9. Write a function that will take this string as Input and return the index of the
//   first character that is non-repeating.

package main

import "fmt"

func nonReapetingFirstCharacter(str string) int {
	for i := 0; i < len(str); i++ {
		isUnique := true
		for j := 0; j < len(str); j++ {
			if i != j && str[i] == str[j] {
				isUnique = false
				break
			}
		}
		if isUnique {
			return i
		}
	}
	return -1
}

// for small time complexity
func firstNonRepeatingCharacter(str string) int {
	frequency := make(map[rune]int)
	for _, char := range str {
		frequency[char]++
	}
	for i, char := range str {
		if frequency[char] == 1 {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(nonReapetingFirstCharacter("mmeles"))
}
