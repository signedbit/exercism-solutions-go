package acronym

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

func Abbreviate(s string) string {
	fields := strings.FieldsFunc(s, isSymbol)
	abbreviation := make([]string, len(fields))
	for i, field := range fields {
		c, _ := utf8.DecodeRuneInString(field)
		abbreviation[i] = string(c)
	}
	return strings.ToUpper(strings.Join(abbreviation, ""))
}

func isSymbol(c rune) bool {
	return !unicode.IsLetter(c) && !unicode.IsNumber(c)
}
