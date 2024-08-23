package ruleparser

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// findRulesTxtURL iterates over the HTML nodes from https://magic.wizards.com/en/rules
// and returns the URL that the TXT link points to (the most recent rules).
func findRulesTxtURL(n *html.Node) string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, attr := range n.Attr {
			if attr.Key == "href" && strings.HasSuffix(attr.Val, ".txt") && strings.HasPrefix(attr.Val, "https://media.wizards.com") {
				return attr.Val
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		foundURL := findRulesTxtURL(c)
		if foundURL != "" {
			return foundURL
		}
	}

	return ""
}

// GetLatestRulesTxtURL downloads the rules webpage and returns the URL of the latest .txt rules file.
// It can fail if it is unable to download the rules webpage, or if if it is unable to find a link to
// a .txt file on the page.
func GetLatestRulesTxtURL() (string, error) {
	resp, err := http.Get("https://magic.wizards.com/en/rules")
	if err != nil {
		return "", fmt.Errorf("failed to download rules webpage from https://magic.wizards.com/en/rules: %w", err)
	}
	defer resp.Body.Close()

	parsed, err := html.Parse(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to parse rules webpage: %w", err)
	}

	foundURL := findRulesTxtURL(parsed)
	if foundURL == "" {
		return "", fmt.Errorf("failed to find rules txt url")
	}

	return foundURL, nil
}
