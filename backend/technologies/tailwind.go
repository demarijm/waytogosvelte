package technologies

import (
	"strings"
)

func isTailwind(bodyString string) bool {
	// Check for some common patterns that might suggest Tailwind CSS.
	tailwindPatterns := []string{
		"text-center",
		"font-bold",
		"bg-blue-500",
		"mt-4",
		"py-2",
		"px-4",
		"w-full",
		"rounded",
		"text-white",
		"hover:bg-blue-700",
		"focus:outline-none",
		"focus:shadow-outline",
		"focus:border-blue-900",
		"border-transparent",
		"border",
		"border-blue-500",
		"hover:border-blue-700",
		"leading-tight",
		"appearance-none",
		"block",
	}

	for _, pattern := range tailwindPatterns {
		if strings.Contains(bodyString, pattern) {
			return true
		}
	}

	return false
}
