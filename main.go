package main

import (
	"fmt"
	"math"
	"unicode"
)

func myAtoi(s string) int {
	// Step 1: Trim leading whitespaces
	i := 0
	n := len(s)
	for i < n && s[i] == ' ' {
		i++
	}

	// Step 2: Handle the sign
	sign := 1
	if i < n && (s[i] == '+' || s[i] == '-') {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}

	// Step 3: Read digits
	num := 0
	for i < n && unicode.IsDigit(rune(s[i])) {
		digit := int(s[i] - '0')
		// Check for overflow before multiplying by 10
		if num > (math.MaxInt32-digit)/10 {
			if sign == 1 {
				return math.MaxInt32
			}
			return math.MinInt32
		}
		num = num*10 + digit
		i++
	}

	// Step 4: Apply the sign
	return num * sign
}

func main() {
	fmt.Println(myAtoi("42"))            // Output: 42
	fmt.Println(myAtoi("   -042"))       // Output: -42
	fmt.Println(myAtoi("1337c0d3"))      // Output: 1337
	fmt.Println(myAtoi("0-1"))           // Output: 0
	fmt.Println(myAtoi("words and 987")) // Output: 0
}
