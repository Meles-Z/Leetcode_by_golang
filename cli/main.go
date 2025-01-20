package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/spf13/cobra"
	"golang.org/x/net/html"
)

// Config represents the JSON configuration file structure
type Config struct {
	Concurrency int    `json:"concurrency"`
	RateLimit   int    `json:"rate_limit"`
	InputFile   string `json:"input_file"`
	OutputFile  string `json:"output_file"`
}

// Result holds the scraping result
type Result struct {
	URL    string
	Title  string
	Status int
	Error  string
}

var (
	config       Config
	retryLimit   = 3 // Maximum retries for failed requests
	totalSuccess int
	totalFailure int
	startTime    time.Time
)

// main is the entry point
func main() {
	var configPath string

	// Cobra command setup
	rootCmd := &cobra.Command{
		Use:   "webscraper",
		Short: "Concurrent Web Scraper with CLI options",
		Run: func(cmd *cobra.Command, args []string) {
			// Load configuration
			err := loadConfig(configPath)
			if err != nil {
				fmt.Println("Error loading config:", err)
				return
			}

			// Start the scraper
			startTime = time.Now()
			err = startScraper()
			if err != nil {
				fmt.Println("Error during scraping:", err)
			}
		},
	}

	// Add CLI flags
	rootCmd.Flags().StringVar(&configPath, "config", "config.json", "Path to the configuration file")
	rootCmd.Execute()
}

// loadConfig reads the configuration file
func loadConfig(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(&config)
}

// startScraper begins the web scraping process
func startScraper() error {
	urls, err := readLines(config.InputFile)
	if err != nil {
		return fmt.Errorf("error reading input file: %v", err)
	}

	// Create output file
	file, err := os.Create(config.OutputFile)
	if err != nil {
		return fmt.Errorf("error creating output file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"URL", "Title", "Status", "Error"})

	urlCh := make(chan string, len(urls))
	resultCh := make(chan Result, len(urls))

	for _, url := range urls {
		urlCh <- url
	}
	close(urlCh)

	limiter := time.NewTicker(time.Second / time.Duration(config.RateLimit))
	defer limiter.Stop()

	var wg sync.WaitGroup

	for i := 0; i < config.Concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for url := range urlCh {
				<-limiter.C
				resultCh <- scrapeWithRetry(url)
			}
		}()
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	for result := range resultCh {
		writer.Write([]string{
			result.URL,
			result.Title,
			fmt.Sprintf("%d", result.Status),
			result.Error,
		})

		if result.Error == "" {
			totalSuccess++
		} else {
			totalFailure++
		}
	}

	printMetrics()
	return nil
}

// scrapeWithRetry attempts to scrape a URL, retrying up to retryLimit times
func scrapeWithRetry(url string) Result {
	var result Result
	for i := 0; i < retryLimit; i++ {
		result = scrapeURL(url)
		if result.Error == "" {
			return result
		}
	}
	return result
}

// scrapeURL fetches the title of the given URL
func scrapeURL(url string) Result {
	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := client.Get(url)
	if err != nil {
		return Result{URL: url, Error: err.Error()}
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Result{URL: url, Status: resp.StatusCode, Error: "Non-OK HTTP status"}
	}

	title, err := extractTitle(resp.Body)
	if err != nil {
		return Result{URL: url, Status: resp.StatusCode, Error: err.Error()}
	}

	return Result{URL: url, Title: title, Status: resp.StatusCode}
}

// extractTitle parses the HTML and retrieves the title
func extractTitle(body io.Reader) (string, error) {
	doc, err := html.Parse(body)
	if err != nil {
		return "", err
	}

	var traverse func(*html.Node) string
	traverse = func(n *html.Node) string {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			return n.FirstChild.Data
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if title := traverse(c); title != "" {
				return title
			}
		}
		return ""
	}

	title := traverse(doc)
	if title == "" {
		return "", fmt.Errorf("title not found")
	}
	return title, nil
}

// readLines reads a file and returns the lines as a slice of strings
func readLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

// printMetrics outputs runtime statistics
func printMetrics() {
	totalTime := time.Since(startTime)
	fmt.Println("Scraping completed!")
	fmt.Printf("Total Time Taken: %v\n", totalTime)
	fmt.Printf("Total Success: %d\n", totalSuccess)
	fmt.Printf("Total Failures: %d\n", totalFailure)
}
