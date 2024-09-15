package main

import (
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		expected  []string
		inputBody string
	}{
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<a href="/path/one">
			<span>Boot.dev</span>
		</a>
		<a href="https://other.com/path/one">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(tc.inputBody, tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			for index, _ := range tc.expected {
				if len(actual) != len(tc.expected) {
					t.Errorf("Test %v - %s FAIL: wrong number of URL(s), expected: %v, actual: %v", i, tc.name, len(tc.expected), len(actual))
				}
				if actual[index] != tc.expected[index] {
					t.Errorf("Test %v - %s FAIL: wrong URL(s), expected: %v, actual: %v", i, tc.name, tc.expected, actual)
				}
			}
		})
	}
}
