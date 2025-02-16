package cmd

import (
	"bufio"
	"fmt"
	"fuzzzz/helper"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	targetURL string
	wordlist  string
)

var rootCmd = &cobra.Command{
	Use:   "fuzzzz",
	Short: "A simple fuzzing tool",
}

var fuzzCmd = &cobra.Command{
	Use:   "fuzzzz",
	Short: "fuzzzz directories on a target URL using a wordlist",
	Run: func(cmd *cobra.Command, args []string) {
		file, err := os.Open(wordlist)
		helper.HandleError(err)
		if err != nil {
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			path := scanner.Text()
			parsedURL, err := url.Parse(targetURL)
			helper.HandleError(err)
			if err != nil {
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

		helper.HandleError(scanner.Err())
	},
}

func Init() {
	rootCmd.AddCommand(fuzzCmd)
	fuzzCmd.Flags().StringVarP(&targetURL, "target", "t", "", "Target URL to scan (required)")
	fuzzCmd.Flags().StringVarP(&wordlist, "wordlist", "w", "", "Path to wordlist file (required)")
	fuzzCmd.MarkFlagRequired("target")
	fuzzCmd.MarkFlagRequired("wordlist")
	fuzzCmd.Execute()
}

func Execute() {
}
