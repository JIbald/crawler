package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(unnormalURL string) (string, error) {
	u, err := url.Parse(unnormalURL)
	if err != nil {
		return "", fmt.Errorf("error parsing url: %w", err)
	}

	result := u.Host + u.Path
	result = strings.ToLower(result)
	result = strings.TrimSuffix(result, "/")
	return result, nil
}
