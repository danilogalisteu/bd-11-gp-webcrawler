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
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: no website provided\nUsage: ./crawler <url>")
		os.Exit(1)
	}
	if len(os.Args) > 2 {
		fmt.Println("Error: too many arguments provided\nUsage: ./crawler <url>")
		os.Exit(1)
	}

	baseURL, err := url.Parse(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	cfg := config{
		pages:   make(map[string]int),
		baseURL: baseURL,
		mu:      &sync.Mutex{},
	}
	fmt.Printf("starting crawl of: %s\n", cfg.baseURL)

	cfg.crawlPage(cfg.baseURL.String())

	fmt.Println("Crawling results:")
	item := 0
	for url, count := range cfg.pages {
		item += 1
		fmt.Printf("%d: %d %s\n", item, count, url)
	}
}
