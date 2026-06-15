package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("couldn't parse URL: %w", err)
	}
	normalizedPath := parsedURL.Host + parsedURL.Path
	normalizedPath = strings.ToLower(normalizedPath)
	normalizedPath = strings.TrimSuffix(normalizedPath, "/")
	return normalizedPath, nil
}
