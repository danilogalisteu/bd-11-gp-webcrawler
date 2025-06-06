package main

import (
	"fmt"
	"log"
	"maps"
	"net/url"
	"os"
	"sort"
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

type ReportItem struct {
	url   string
	count int
}

func sortPages(pages map[string]int) []ReportItem {
	sorted := []ReportItem{}
	for k, v := range maps.All(pages) {
		sorted = append(sorted, ReportItem{url: k, count: v})
	}
	sort.Slice(sorted, func(i, j int) bool {return sorted[i].url < sorted[j].url})
	sort.SliceStable(sorted, func(i, j int) bool {return sorted[i].count > sorted[j].count})
	return sorted
}

func printReport(pages map[string]int, baseURL string) {
	fmt.Printf("=============================\n")
	fmt.Printf("  REPORT for %s\n", baseURL)
	fmt.Printf("=============================\n")
	for _, item := range sortPages(pages) {
		fmt.Printf("Found %d internal links to %s\n", item.count, item.url)
	}
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
	printReport(cfg.pages, cfg.baseURL.String())
}
