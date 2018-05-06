package luhn

import (
	"regexp"
	"unicode"
)

func Valid(checksum string) bool {
	checksum = regexp.MustCompile("\\s").ReplaceAllString(checksum, "")

	if len(checksum) <= 1 {
		return false
	}

	sum := 0
	chars := []rune(checksum)
	for i, length := 0, len(chars); i < length; i++ {
		char := chars[length-i-1]
		if !unicode.IsDigit(char) {
			return false
		}
		digit := int(char - '0')
		if i%2 == 1 {
			digit *= 2
			if digit >= 10 {
				digit -= 9
			}
		}
		sum += digit
	}

	return sum%10 == 0
}
