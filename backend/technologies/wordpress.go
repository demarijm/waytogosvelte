package technologies

import (
	"strings"
)

func isWordpress(bodyString string) bool {
	// List of WordPress-specific identifiers to check
	wordpressIdentifiers := []string{
		"/wp-content",
		"/wp-includes",
		"/wp-admin",
		"wp-embed.min.js",         // common WordPress script
		"wp-emoji-release.min.js", // another common WordPress script
	}

	// Convert the bodyString to lowercase to ensure the check isn't case sensitive
	bodyString = strings.ToLower(bodyString)

	for _, identifier := range wordpressIdentifiers {
		if strings.Contains(bodyString, identifier) {
			return true
		}
	}

	// If none of the WordPress-specific identifiers exist, return false
	return false
}
