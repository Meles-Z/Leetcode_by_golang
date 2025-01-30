// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"os"
// )

// type Skills struct {
// 	LanguagesKnown []string
// 	Passion        []string
// }

// type UserData struct {
// 	UserId          int
// 	UserName        string
// 	IsActive        bool
// 	Roles           []string
// 	TrainingCreadit float64 `json:"training-credit"`
// 	Skills          Skills
// }

// func main() {
// 	var userInfo UserData

// 	file, err := os.ReadFile("test.json")
// 	if err != nil {
// 		log.Fatalf("Error to read json file:%s", err)
// 	}
// 	err = json.Unmarshal(file, &userInfo)
// 	if err != nil {
// 		log.Fatal("Error to atach file to user info", err)
// 	}
// 	fmt.Println(userInfo)
// }

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"os"
// )

// type Skills struct {
// 	LanguagesKnown []string
// 	Passion        []string
// }
// type User struct {
// 	UserId          int
// 	UserName        string
// 	IsActive        bool
// 	Roles           []string
// 	TrainingCreadit float64 `json:"training-credit"`
// 	Skills          Skills
// }

// func main() {
// 	var user []User
// 	file, err := os.ReadFile("main.json")
// 	if err != nil {
// 		log.Fatalf("Error to read main.json file:%s", err)
// 	}

// 	err = json.Unmarshal(file, &user)
// 	if err != nil {
// 		log.Fatal("Error to unmarshaling:", err)
// 	}
// 	fmt.Println(user)

// }

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type Book struct {
// 	Id     int    `json:"id"`
// 	Title  string `json:"title"`  // Keep "title" lowercase
// 	Author string `json:"author"`
// }

// func main() {
// 	b := Book{
// 		Id:     1,
// 		Title:  "The rich man in Bayblon",
// 		Author: "Dr I do not know",
// 	}

// 	goOb := map[string]int{
// 		"Meles": 1,
// 		"Tola":  2,
// 		"Fufa":  3,
// 	}
// 	bytes, err := json.Marshal(goOb)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(string(bytes))

// 	bb, err := json.Marshal(b)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(string(bb))
// }

// package main

// import (
// 	"bufio"
// 	"encoding/json"
// 	"fmt"
// 	"os"
// )

// type Seller struct {
// 	Id          int
// 	Name        string
// 	CountryCode string
// }

// type Product struct {
// 	Id     int
// 	Name   string
// 	Seller Seller
// 	Price  float64
// }

// func main() {

// 	product := []Product{
// 		Product{
// 			Id:   1,
// 			Name: "Vegt",
// 			Seller: Seller{
// 				Id:          1,
// 				Name:        "fff",
// 				CountryCode: "+251",
// 			},
// 			Price: 50.5,
// 		},
// 		Product{
// 			Id:   2,
// 			Name: "Vv",
// 			Seller: Seller{
// 				Id:          2,
// 				Name:        "f",
// 				CountryCode: "54",
// 			},
// 			Price: 40.5,
// 		},
// 		Product{
// 			Id:   3,
// 			Name: "fft",
// 			Seller: Seller{
// 				Id:          3,
// 				Name:        "g",
// 				CountryCode: "4334",
// 			},
// 			Price: 40.4,
// 		},
// 	}

// 	prod, _ := json.MarshalIndent(product, "", " ")
// 	fmt.Println(string(prod))

// 	file, err := os.Create("ff.txt")
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	lls := bufio.NewWriter(file)
// 	lls.Write(prod)
// 	lls.Flush()
// }

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Seller struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	CountryCode string `json:"country_code,omitempty"`
}
type Product struct {
	Id     int    `json:"-"`
	Name   string `json:"name"`
	Seller Seller `json:"seller"`
	Price  int    `json:"price"`
}

func main() {
	fmt.Println("Let us dive into json unmarshaller")
	// its totall inverse of json.Marshal
	// takes butes

	file, err := os.ReadFile("ff.txt")
	if err != nil {
		log.Fatalln(err)
	}
	var prod Product

	// err := json.Unmarshal(file, &prod)
	
	// fmt.Println(mrsh)

}
