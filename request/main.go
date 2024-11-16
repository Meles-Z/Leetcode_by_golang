package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Post struct {
	ID     int    `json:"id"`
	UserId int    `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/2")
	if err != nil {
		fmt.Println("Error to reading request :", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error to reading response body: ", err)
		return
	}
	var post Post
	err = json.Unmarshal(body, &post)
	if err != nil {
		fmt.Println("Error to unmarshalling json: ", post)
		return
	}
	fmt.Println("ID:", post.ID)
	fmt.Println("UserId:", post.UserId)
	fmt.Println("Title:", post.Title)
	fmt.Println("Body:", post.Body)
}
