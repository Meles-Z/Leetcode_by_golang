package main

import "fmt"

type Family struct {
	person []Person
}
type Person struct {
	Name          string
	Age           int
	MartialStatus string
}

func defaultPerson() Person {
	return Person{
		Name:          "Meles",
		Age:           12,
		MartialStatus: "married",
	}
}

func NewFamily() *Family {
	defaultPerson := defaultPerson()
	return &Family{
		person: []Person{defaultPerson},
	}
}

func (family *Family) AddFamily(name string, age int, martial string) {
	person := Person{
		Name:          name,
		Age:           age,
		MartialStatus: martial,
	}
	family.person = append(family.person, person)
}

func main() {
	family := NewFamily()
	family.AddFamily("namaste", 23, "rrt")
	for _, val:=range family.person{
		fmt.Println(val)
	}
}
