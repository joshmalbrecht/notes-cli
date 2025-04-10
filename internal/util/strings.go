package util

import "regexp"

func ReplaceSpaces(s string, char string) string {
	regex := regexp.MustCompile(`\s+`)
	return regex.ReplaceAllString(s, char)
}
