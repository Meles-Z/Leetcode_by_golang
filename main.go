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
	"fmt"
	"strconv"
)


 
func main(){
	ch:=make(chan string)
	for i:=0; i<10; i++{
		go func(){
			for j:=0; j<10; j++{
				ch<-"Gorutines: "+strconv.Itoa(i)
			}
		}()
	}
	for k:=0; k<100; k++{
		fmt.Println(<-ch)
	}
}
