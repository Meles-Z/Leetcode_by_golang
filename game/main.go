package main

import (
	"fmt"
)

func GuessNumber(num int, secret int) {
	if num == secret {
		fmt.Println("Congratulations! You are the winner!")
		return
	}
	if num > secret {
		fmt.Printf("Hint: %d is greater than the secret number!\n", num)
	} else {
		fmt.Printf("Hint: %d is less than the secret number!\n", num)
	}
}

func main() {
	secret := 23
	var num int
	fmt.Println("Enter your guess:")
	for {
		fmt.Scanln(&num)
		GuessNumber(num, secret)
		if num == secret {
			break
		}
	}
}
