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
}
