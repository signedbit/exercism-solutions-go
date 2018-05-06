package isogram

import (
	"regexp"
	"strings"
)

func IsIsogram(word string) bool {
	nonLetters := regexp.MustCompile("[^a-z]")
	characters := []rune(nonLetters.ReplaceAllString(strings.ToLower(word), ""))
	set := make(map[rune]bool, len(word))

	for _, c := range characters {
		if set[c] {
			return false
		}
		set[c] = true
	}

	return true
}
