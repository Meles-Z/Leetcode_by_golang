package main

import "fmt"

func Login(user, pass string) {
	if user == "username" && pass == "password" {
		fmt.Println("Login succesfull")
	} else {
		fmt.Println("Incorrect username or password")
	}
}
func Seach(s string, target byte) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == target {
			return true
		}
	}
	return false
}
func main() {
	Login("username", "password")
	Login("username", "passw")
	fmt.Println(Seach("Generator", 'G'))
}
