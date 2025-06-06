package main

import (
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	res := make([]string, 0)

	doc, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return res, err
	}

	for node := range doc.Descendants() {
		if node.Type == html.ElementNode && node.Data == "a" {
			for _, attr := range node.Attr {
				if attr.Key == "href" {
					link := attr.Val
					if strings.HasPrefix(link, "/") {
						link = strings.TrimRight(rawBaseURL, "/") + link
					}
					res = append(res, link)
				}
			}
		}
	}

	return res, nil
}
