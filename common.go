package e2g_utils

import "strings"

func matchSuffix(suffixes []string, fileName string) bool {
	for _, suffix := range suffixes {
		if strings.HasSuffix(fileName, suffix) {
			return true
		}
	}
	return false
}
