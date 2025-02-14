package main

import (
	"bufio"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <target URL> <wordlist file>")
		return
	}
	target := strings.TrimSuffix(os.Args[1], "/")
	wordlist := os.Args[2]

	file, err := os.Open(wordlist)
	if err != nil {
		fmt.Printf("Error opening wordlist: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		path := scanner.Text()
		//url := fmt.Sprintf("%s/%s", target, path)
	
		parsedURL, err := url.Parse(target)
		if err != nil {
			fmt.Printf("Error parsing URL: %v\n", err)
			return
		}
		parsedURL.Path = strings.TrimSuffix(parsedURL.Path, "/") + "/" + strings.TrimPrefix(path, "/")
		url := parsedURL.String()
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Error requesting %s: %v\n", url, err)
			continue
		}
		fmt.Printf("[%d] %s\n", resp.StatusCode, url)
		resp.Body.Close()
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading wordlist: %v\n", err)
	}
}
