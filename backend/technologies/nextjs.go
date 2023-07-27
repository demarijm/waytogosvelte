package technologies

import (
	"strings"
)

func isNextJS(bodyString string) bool {
	// Check for the __NEXT_DATA__ pattern that suggests Next.js
	if strings.Contains(bodyString, "__NEXT_DATA__") ||
		strings.Contains(bodyString, "/_next/static/") ||
		strings.Contains(bodyString, "next-route-announcer") {
		return true
	}

	return false
}
