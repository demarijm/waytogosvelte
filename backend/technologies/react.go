package technologies

import (
	"strings"
)

func isReact(bodyString string) bool {
	// Check for some common patterns that might suggest React.
	reactPatterns := []string{
		"__NEXT_DATA__",
		"div id=\"__next\"",
		"data-reactroot",
		"data-reactid",
		"data-react-checksum",
		"for=\"react",
		"id=\"react",
		"class=\"react",
		"\"react-root\"",
	}

	for _, pattern := range reactPatterns {
		if strings.Contains(strings.ToLower(bodyString), pattern) {
			return true
		}
	}

	return false
}
