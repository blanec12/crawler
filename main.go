package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	if len(args) < 3 {
		fmt.Println("not enough arguments provided")
		fmt.Println("usage: crawler <baseURL> <maxConcurrency> <maxPages>")
		os.Exit(1)
	}

	if len(args) > 3 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	rawBaseURL := args[0]
	maxConcurrencyString := args[1]
	maxPagesString := args[2]

	maxConcurrency, err := strconv.Atoi(maxConcurrencyString)
	if err != nil {
		fmt.Printf("couldn't convert maxConcurrencyString to int: %v", err)
		return
	}

	maxPages, err := strconv.Atoi(maxPagesString)
	if err != nil {
		fmt.Printf("couldn't convert maxPagesString to int: %v", err)
		return
	}

	cfg, err := configure(rawBaseURL, maxConcurrency, maxPages)
	if err != nil {
		fmt.Printf("couldn't get configuration: %v", err)
		return
	}

	cfg.wg.Add(1)
	go cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()

	for url := range cfg.pages {
		fmt.Printf("found: %s\n", url)
	}

	writeJSONReport(cfg.pages, "report.json")
}
