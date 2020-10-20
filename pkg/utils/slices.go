package utils

import (
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func UintInSlice(elem uint, list []uint) bool {
	for _, b := range list {
		if b == elem {
			return true
		}
	}
	return false
}

// TrimAllSpaces returns the string passed as argument without spaces and in lower key
func TrimAllSpaces(value string) string {
	value = strings.TrimSpace(value)
	space := regexp.MustCompile(`\s+`)
	value = space.ReplaceAllString(value, " ")
	return value
}

// ToSnakeCase returns the string in snake case format
func ToSnakeCase(value string) string {
	value = strings.ToLower(value)
	value = strings.ReplaceAll(value, " ", "_")
	return value
}

// RemoveBracketsContent returns the string without brackets and without the content inside of them
func RemoveBracketsContent(value string) string {
	value = strings.Split(value, " (")[0]
	return value
}

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

// PrepareString removes special chars, words between brackets, alternative text and covert the string to lower case
func PrepareString(value string) string {
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	result, _, _ := transform.String(t, value)
	result = strings.ToLower(result)
	result = strings.Split(result, " /")[0]
	result = strings.Split(result, " (")[0]
	result = strings.ReplaceAll(result, "-", " ")
	result = TrimAllSpaces(result)
	return result
}
