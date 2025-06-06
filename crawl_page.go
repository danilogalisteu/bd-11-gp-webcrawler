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

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
    cfg.mu.Lock()
    defer cfg.mu.Unlock()
	count, exists := cfg.pages[normalizedURL]
	if exists {
		cfg.pages[normalizedURL] = count + 1
		return false
	}
	cfg.pages[normalizedURL] = 1
	return true
}

func (cfg *config) crawlPage(rawCurrentURL string) {
	defer cfg.wg.Done()

	sameDomain, err := checkSameDomain(cfg.baseURL.String(), rawCurrentURL)
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

	if !cfg.addPageVisit(normUrl) {
		return
	}

	fmt.Printf("crawling %s...\n", rawCurrentURL)
	html, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("%s: non-HTML link found, skipping...\n", rawCurrentURL)
	}

	urls, err := getURLsFromHTML(html, cfg.baseURL.String())
	if err != nil {
		log.Fatal(err)
	}
	for _, url := range urls {
		fmt.Printf("checking %s...\n", url)
		cfg.wg.Add(1)
		go cfg.crawlPage(url)
	}
}
