package utils

import (
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func NormalizeSpaces(value string) string {

	if strings.HasPrefix(value, " ") || strings.HasSuffix(value, " ") {
		value = strings.TrimSpace(value)
	}
	if strings.Contains(value, "  ") {
		values := strings.Fields(value)
		value = strings.Join(values, " ")
	}
	return value
}

func RemoveAccentuation(value string, lowerCase bool) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	result, _, err := transform.String(t, value)
	if err != nil {
		result = value
	}

	if lowerCase {
		result = strings.ToLower(result)
	}

	return result
}
