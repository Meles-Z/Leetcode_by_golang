package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func webScrapper(ctx context.Context, url string, result chan<- string) {
	client := http.DefaultClient
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		result <- fmt.Sprintf("Error to create reques:%s", err)
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		result <- fmt.Sprintf("Error to send request:%s", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		result <- fmt.Sprintf("Error to read body file:%s", err)
		return
	}
	result <- fmt.Sprintf("The result of request body:%s", string(body))
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	urls := []string{
		"https://api.example.com/users",
		"https://api.example.com/products",
		"https://api.example.com/orders",
	}
	ch := make(chan string)
	for _, url := range urls {
		go webScrapper(ctx, url, ch)
	}
	for range urls {
		fmt.Println(<-ch)
	}
}
