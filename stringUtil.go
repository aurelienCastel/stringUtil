package stringUtil

import "strings"

func HasOneOfSuffixes(str string, suffixes []string) bool {
	for _, suffix := range suffixes {
		if strings.HasSuffix(str, suffix) {
			return true
		}
	}
	return false
}

func FileExtension(str string) string {
	for i := len(str) - 1; i > -1; i-- {
		if str[i] == '.' {
			return str[i+1 : len(str)]
		}
	}
	return ""
}
