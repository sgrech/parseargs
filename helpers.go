package parseargs

import (
	"strings"
)

func stripDashes(s string) string {
	return strings.ReplaceAll(s, "-", "")
}

func parseLongOption(s string) (option string, value string) {
	if strings.Contains(s, "=") {
		kvp := strings.Split(s, "=")
		return kvp[0], kvp[1]
	}
	return s, ""
}

func parseShortOption(s string) (options []string) {
	if len(s) > 1 {
		return strings.Split(s, "")
	}
	return []string{s}
}
