package main

import (
	"fmt"
	"io"
	"net/http"
)

func getHTML(rawURL string) (string, error) {
	res, err := http.Get(rawURL)
	if err != nil {
		return "", err
	}

	if res.StatusCode > 399 {
		return "", fmt.Errorf("http.Get for '%s' failed with status: %d", rawURL, res.StatusCode)
	}

	if res.Header.Get("Content-Type") != "text/html" {
		return "", fmt.Errorf("content-type for '%s' is not 'text/html'", rawURL)
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return "", err
	}

	return string(body), nil
}
