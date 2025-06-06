package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"sync"
)

type config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	wg                 *sync.WaitGroup
	concurrencyControl chan struct{}
	maxPages           int
}

func main() {
	maxConcurrency := 5
	maxPages := 100

	if len(os.Args) < 2 {
		fmt.Println("Error: no website provided\nUsage: ./crawler <url> [maxConcurrency [maxPages]]")
		os.Exit(1)
	}
	if len(os.Args) > 4 {
		fmt.Println("Error: too many arguments provided\nUsage: ./crawler <url>")
		os.Exit(1)
	}
	if len(os.Args) > 2 {
		_, err := fmt.Sscanf(os.Args[2], "%d", &maxConcurrency)
		if err != nil {
			log.Fatal(err)
		}
		if len(os.Args) > 3 {
			_, err := fmt.Sscanf(os.Args[3], "%d", &maxPages)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	baseURL, err := url.Parse(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	cfg := config{
		pages:              make(map[string]int),
		baseURL:            baseURL,
		mu:                 &sync.Mutex{},
		wg:                 &sync.WaitGroup{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		maxPages:           maxPages,
	}
	fmt.Printf("starting crawl of: %s\n", cfg.baseURL)

	cfg.wg.Add(1)
	go cfg.crawlPage(cfg.baseURL.String())

	cfg.wg.Wait()
	fmt.Println("Crawling results:")
	item := 0
	for url, count := range cfg.pages {
		item += 1
		fmt.Printf("%d: %d %s\n", item, count, url)
	}
}
