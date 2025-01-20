package main

import "fmt"

type customeFunc func(*Person)
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


func NewFamily(fam ...customeFunc) *Family {
	defaultPerson := defaultPerson()
	for _, person := range fam {
		person(&defaultPerson)
	}
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

func UpdateName(name string) customeFunc {
	return func(p *Person) {
		p.Name = name
	}
}

func main() {
	family := NewFamily(UpdateName("fufa"))
	family.AddFamily("namaste", 23, "rrt")

	for _, val := range family.person {
		fmt.Println(val)
	}
}
