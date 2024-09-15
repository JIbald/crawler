package main

import (
	"fmt"
	"golang.org/x/net/html"
	"strings"
)

func getHrefValuesFromHTML(tree *html.Node) []string {
	result := []string{}
	if tree.Type == html.ElementNode && tree.Data == "a" {
		for _, anchor := range tree.Attr {
			if anchor.Key == "href" {
				result = append(result, anchor.Val)
				break
			}
		}
	}
	for child := tree.FirstChild; child != nil; child = child.NextSibling {
		result = append(result, getHrefValuesFromHTML(child)...)
	}
	return result
}

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	result := []string{}
	tree, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return []string{}, fmt.Errorf("error Parsing ioReader: %w", err)
	}
	result = getHrefValuesFromHTML(tree)
	fmt.Println(result)

	return result, nil
}
