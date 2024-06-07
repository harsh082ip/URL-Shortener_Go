package helpers

import (
	"fmt"
	"net/http"
	"net/url"
)

// IsValidURL checks if the URL is valid and reachable
func IsValidURL(rawURL string) bool {
	// Parse the URL to check if it's valid
	parsedURL, err := url.ParseRequestURI(rawURL)
	if err != nil {
		fmt.Printf("Invalid URL: %v\n", err)
		return false
	}

	// Make a GET request to the URL to check if it's reachable
	resp, err := http.Get(parsedURL.String())
	if err != nil {
		fmt.Printf("Error reaching URL: %v\n", err)
		return false
	}
	defer resp.Body.Close()

	// Check if the status code indicates the URL is reachable
	if resp.StatusCode == http.StatusOK {
		return true
	}

	fmt.Printf("URL returned status code: %d\n", resp.StatusCode)
	return false
}
