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

func pop(slice *[]uint64) (value uint64, is_empty bool) {
	if len(*slice) == 0 {
		return 0, true // Empty slice
	}

	lastIndex := len(*slice) - 1
	value = (*slice)[lastIndex]
	*slice = (*slice)[:lastIndex]

	return value, false
}

func validateBrackets(input string) bool {
	depth := 0
	for _, ch := range input {
		if ch == '[' {
			depth++
		} else if ch == ']' {
			depth--
			if depth < 0 {
				return false
			}
		}
	}
	return depth == 0
}
