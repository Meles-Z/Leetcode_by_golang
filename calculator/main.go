package main

import (
	"fmt"
)

// add two number
func add(x, y int) int {
	return x + y
}

// substruct two number
func sub(x, y int) int {
	return x - y
}

// devide two number

func div(x, y int) int {
	if y == 0 {
		fmt.Println("Number is not divide by zero!")
		return 0
	}
	return x / y
}

// multiply two number
func mult(x, y int) int {
	return x * y
}
func main() {
	var num1, num2, choice int

	for {
		fmt.Println("\nWelcome to Scientific Calculator!")
		fmt.Println("Enter your option:")
		fmt.Println("1.Add")
		fmt.Println("2.Substruct")
		fmt.Println("3.Devide")
		fmt.Println("4.Multiplication")
		fmt.Println("5.Exit")
		fmt.Println("Enter your choice:")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Print("Enter number 1:")
			fmt.Scan(&num1)
			fmt.Println("Enter number 2")
			fmt.Scan(&num2)
			fmt.Printf("The sum of %d + %d = %d\n", num1, num2, add(num1, num2))
		case 2:
			fmt.Print("Enter number 1:")
			fmt.Scan(&num1)
			fmt.Println("Enter number 2")
			fmt.Scan(&num2)
			fmt.Printf("The sub of %d - %d = %d\n", num1, num2, sub(num1, num2))

		case 3:
			fmt.Print("Enter number 1:")
			fmt.Scan(&num1)
			fmt.Println("Enter number 2")
			fmt.Scan(&num2)
			fmt.Printf("The div of %d / %d = %d\n", num1, num2, div(num1, num2))

		case 4:
			fmt.Print("Enter number 1:")
			fmt.Scan(&num1)
			fmt.Println("Enter number 2")
			fmt.Scan(&num2)
			fmt.Printf("The mult of %d * %d = %d\n", num1, num2, mult(num1, num2))

		case 5:
			fmt.Println("Goodbyte rest in peace bruv!")
			return
		default:
			fmt.Println("Icorrect choice please try it again!")
		}
	}
}
