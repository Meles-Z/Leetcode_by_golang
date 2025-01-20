package main

import (
	"fmt"
	"os"
	"strconv"
)

func sum(num1, num2 int) int {
	return num1 + num2
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Your input need like this format: go run main.go <num1> <num2>")
		return
	}

	arg1 := os.Args[1]
	arg2 := os.Args[2]

	num1, err := strconv.Atoi(arg1)
	if err != nil {
		fmt.Println(err)
		return
	}
	num2, err := strconv.Atoi(arg2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("The sum of %[1]d + %[2]d = %d\n", num1, num2, sum(num1, num2))
}
