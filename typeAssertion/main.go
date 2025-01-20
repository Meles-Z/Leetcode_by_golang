package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func getStudent(name interface{}, age interface{}) Student {
	var stud Student
	stud.Name = name.(string)
	stud.Age = age.(int)
	return stud
}
func main() {
	fmt.Println("Hello test!")
	var name interface{} = "Meles"
	var age interface{} = 25
	daynamicStud := getStudent(name, age)
	fmt.Println(daynamicStud.Name)
	fmt.Println(daynamicStud.Age)

}
