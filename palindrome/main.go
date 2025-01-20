package main

import "fmt"

func palindome(n int) bool {
	numString := fmt.Sprintf("%d", n)
	if len(numString) == 0 {
		fmt.Println("Number is empty")
	}

	for i := 0; i < len(numString); i++ {
		if numString[i] != numString[len(numString)-1-i] {
			return false
		}
	}
	return true
}
func main() {
	fmt.Println(palindome(1344))

}