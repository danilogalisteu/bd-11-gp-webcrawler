package main

import (
	"fmt"
	"log"
	"net/url"
)

func checkSameDomain(baseURL, currentURL string) (bool, error) {
	partsBase, err := url.Parse(baseURL)
	if err != nil {
		return false, err
	}
	partsCurrent, err := url.Parse(currentURL)
	if err != nil {
		return false, err
	}
	return partsBase.Hostname() == partsCurrent.Hostname(), nil
}

func addPageVisit(normalizedURL string, pages map[string]int) (isFirst bool) {
	count, exists := pages[normalizedURL]
	if exists {
		pages[normalizedURL] = count + 1
		return false
	}
	pages[normalizedURL] = 1
	return true
}

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	sameDomain, err := checkSameDomain(rawBaseURL, rawCurrentURL)
	if err != nil {
		log.Fatal(err)
	}
	if !sameDomain {
		return
	}

	normUrl, err := normalizeURL(rawCurrentURL)
	if err != nil {
		log.Fatal(err)
	}

	if !addPageVisit(normUrl, pages) {
		return
	}

	fmt.Printf("crawling %s...\n", rawCurrentURL)
	html, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("%s: non-HTML link found, skipping...\n", rawCurrentURL)
	}

	urls, err := getURLsFromHTML(html, rawBaseURL)
	if err != nil {
		log.Fatal(err)
	}
	for _, url := range urls {
		fmt.Printf("checking %s...\n", url)
		crawlPage(rawBaseURL, url, pages)
	}
}
