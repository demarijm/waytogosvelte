package technologies

import (
	"strings"
)

func isWix(bodyString string) bool {
	// List of Wix-specific identifiers to check
	wixIdentifiers := []string{
		"<!--wix",
		"wix-image",
		".wix.com",
		"wixapps.net",
	}

	// Convert the bodyString to lowercase to ensure the check isn't case sensitive
	bodyString = strings.ToLower(bodyString)

	for _, identifier := range wixIdentifiers {
		if strings.Contains(bodyString, identifier) {
			return true
		}
	}

	// If none of the Wix-specific identifiers exist, return false
	return false
}
