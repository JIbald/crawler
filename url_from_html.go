package main

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	result := []string{}
	result = append(result, htmlBody)
	result = append(result, rawBaseURL)

	return result, nil
}
