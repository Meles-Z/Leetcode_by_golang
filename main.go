/*
package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func Scrapper(cxt context.Context, url string, result chan<- string) {
	// send request to get response
	req, err := http.NewRequestWithContext(cxt, http.MethodGet, url, nil)
	if err != nil {
		result <- fmt.Sprintf("Error feaching:%s:%s", url, err)
		return
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		select {
		case result <- fmt.Sprintf("Error creating request for:%s:%s", url, err):
		case <-cxt.Done():
			result <- fmt.Sprintf("Request to :%s timeout", url)
		}
		return
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		result <- fmt.Sprintf("Error feaching:%s:%s", url, err)
		return
	}
	result <- fmt.Sprintf("Response from:%s:\n %s", url, string(body))
}
func main() {
	start := time.Now()
	urls := []string{
		"https://jsonplaceholder.org/users",
		"https://example.com",
		// "https://google.com",
	}
	ch := make(chan string)
	for _, url := range urls {
		go func(url string) {
			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()
			Scrapper(ctx, url, ch)
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-ch)
	}
	close(ch)
	elipse := time.Since(start)
	fmt.Println("The response time is:", elipse)
}
*/

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Address  struct {
		Street  string `json:"street"`
		Suite   string `json:"suite"`
		City    string `json:"city"`
		Zipcode string `json:"zipcode"`
	} `json:"address"`
	Phone   string `json:"phone"`
	Website string `json:"website"`
	Company struct {
		Name        string `json:"name"`
		CatchPhrase string `json:"catchPhrase"`
		Bs          string `json:"bs"`
	} `json:"company"`
}



func WebScrapper(url string) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("error occured:%s", err)
		return
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	var users []User
	err = json.Unmarshal(body, &users)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, user := range users {
		fmt.Println(user.Email)
	}
}

func main() {
	url := "https://jsonplaceholder.typicode.com/users"
	WebScrapper(url)
}
