package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: no website provided\nUsage: ./crawler <url>")
		os.Exit(1)
	}
	if len(os.Args) > 2 {
		fmt.Println("Error: too many arguments provided\nUsage: ./crawler <url>")
		os.Exit(1)
	}

	baseUrl := os.Args[1]
	fmt.Printf("starting crawl of: %s\n", baseUrl)

	pages := make(map[string]int)
	crawlPage(baseUrl, baseUrl, pages)

	fmt.Println("Crawling results:")
	item := 0
	for url, count := range pages {
		item += 1
		fmt.Printf("%d: %d %s\n", item, count, url)
	}
}
