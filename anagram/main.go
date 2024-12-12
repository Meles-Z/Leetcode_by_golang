package main

import (
	"fmt"
)

func frequencyOfCharacter(s string) map[string]int {
	memo := make(map[string]int)
	for _, char := range s {
		memo[string(char)]++
	}
	return memo
}

func findFirstCharacterNonReate(s string) rune {
	memo := make(map[rune]int)
	// calculate all the frequencey
	for _, char := range s {
		memo[char]++
	}
	// filter the frequnecy only one
	for _, char := range s {
		if memo[char] == 1 {
			return char
		}
	}
	return '_'
}

func anagrams(s string, p string) bool {
	if len(s) != len(p) {
		return false
	}
	mapS := make(map[rune]int)
	mapT := make(map[rune]int)
	for _, char := range s {
		mapS[char]++
	}
	for _, char := range p {
		mapT[char]++
	}

	for i, value := range mapS {
		if mapT[i] != value {
			return false
		}
	}
	return true

}
func main() {
	fmt.Println(anagrams("meles", "seler"))
}
