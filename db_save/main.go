package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Website  string `json:"website"`
}

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	host := "localhost"
	port := 5432
	dbName := "test"
	usrName := "meles"
	password := "meles@123"
	dns := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable", host, port, dbName, usrName, password)

	testDb, err := gorm.Open(postgres.Open(dns), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	DB = testDb
	err = DB.AutoMigrate(&User{})
	if err != nil {
		return nil, err
	}
	fmt.Println("Database connected successfully")
	return testDb, nil
}

func WebScrapper(url string) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error occurred during request: %s\n", err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error occurred while reading response: %s\n", err)
		return
	}

	var users []User
	err = json.Unmarshal(body, &users)
	if err != nil {
		fmt.Printf("Error occurred during JSON unmarshaling: %s\n", err)
		return
	}

	for _, user := range users {
		result := DB.Create(&user)
		if result.Error != nil {
			fmt.Printf("Failed to insert user: %s\n", result.Error)
		} else {
			fmt.Printf("Inserted user: %s\n", user.Email)
		}
	}
}

func main() {
	db, err := InitDB()
	if err != nil {
		fmt.Println("Failed to initialize database:", err)
		return
	}
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	url := "https://jsonplaceholder.typicode.com/users"
	WebScrapper(url)
}
