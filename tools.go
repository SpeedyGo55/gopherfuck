package main

import "strings"

func filterString(s string, allowed string) string {
	// Create a map of allowed characters for O(1) lookup
	allowedChars := make(map[rune]bool)
	for _, ch := range allowed {
		allowedChars[ch] = true
	}

	// Build the filtered string
	var result strings.Builder
	result.Grow(len(s)) // Pre-allocate to avoid reallocation

	for _, ch := range s {
		if allowedChars[ch] {
			result.WriteRune(ch)
		}
	}

	return result.String()
}
