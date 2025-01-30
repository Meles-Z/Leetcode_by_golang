package main

import "fmt"

type TestStruct struct {
	Name    string
	Email   string
	Address struct {
		City   string
		Post   string
		Street string
	}
}

type OptsTest struct {
	test []TestStruct
}

func (t *OptsTest) addTest(name, email string) {
	testStr := TestStruct{
		Name:  name,
		Email: email,
		Address: struct{City string; Post string; Street string}{
			City: "AA",
			Post: "123",
			Street: "Shshs",
		},
	}
	t.test = append(t.test, testStr)
}

func main() {
	test := OptsTest{}
	test.addTest("F", "f@gmail.com")
	test.addTest("T", "t@gmail.com")
	test.addTest("F", "f@gmail.com")
	test.addTest("T", "t@gmail.com")
	for _, v := range test.test {
		fmt.Println(v)
	}

	app:=struct{
		version string
		lastUpdate string
	}{
		version: "v1.1",
		lastUpdate: "jan 12, 2012",
	}
	fmt.Println(app)
}
