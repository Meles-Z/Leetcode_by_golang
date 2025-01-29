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

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Skills struct {
	LanguagesKnown []string
	Passion        []string
}
type User struct {
	UserId          int
	UserName        string
	IsActive        bool
	Roles           []string
	TrainingCreadit float64 `json:"training-credit"`
	Skills          Skills
}

func main() {
	var user []User
	file, err := os.ReadFile("main.json")
	if err != nil {
		log.Fatalf("Error to read main.json file:%s", err)
	}

	err = json.Unmarshal(file, &user)
	if err != nil {
		log.Fatal("Error to unmarshaling:", err)
	}
	fmt.Println(user)

}
