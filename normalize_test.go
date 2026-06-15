package main

import "testing"

func TestNormalizedURL(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "remove scheme",
			input:    "https://crawler-test.com/path",
			expected: "crawler-test.com/path",
		},
		{
			name:     "remove trailing slash",
			input:    "https://crawler-test.com/path/",
			expected: "crawler-test.com/path",
		},
		{
			name:     "lowercase capital letters",
			input:    "https://CRAWLER-TEST.com/PATH",
			expected: "crawler-test.com/path",
		},
		{
			name:     "remove scheme and capitals and trailing slash",
			input:    "http://CRAWLER-TEST.com/path/",
			expected: "crawler-test.com/path",
		},
	}

	for i, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := normalizeURL(test.input)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexepected error: %v", i, test.name, err)
			}
			if actual != test.expected {
				t.Errorf("Test %v - '%s' FAIL: expected URL: %v, actual: %v", i, test.name, test.expected, actual)
			}
		})
	}
}
