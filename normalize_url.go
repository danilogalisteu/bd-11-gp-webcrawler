package main

import (
	"net/url"
	"strings"
)

func normalizeURL(rawURL string) (string, error) {
	parts, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}
	normUrl := strings.TrimRight(parts.Host + parts.Path, "/")
	return normUrl, nil
}
