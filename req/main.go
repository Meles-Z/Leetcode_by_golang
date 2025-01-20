package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	user := struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}{
		Name:     "Meles",
		Email:    "meles@gmail.com",
		Password: "password",
	}
	requestBody, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}
	stringBytes := bytes.NewBuffer(requestBody)
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://www.jsonplaceholder.org/", stringBytes)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Request body error:", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
}