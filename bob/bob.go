package bob

import (
	"regexp"
	"strings"
)

func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	switch {
	case isForcefulQuestion(remark):
		return "Calm down, I know what I'm doing!"
	case isShouting(remark):
		return "Whoa, chill out!"
	case isAskingAQuestion(remark):
		return "Sure."
	case isSilence(remark):
		return "Fine. Be that way!"
	}
	return "Whatever."
}

func isQuestion(remark string) bool {
	return strings.HasSuffix(remark, "?")
}

func isShouting(remark string) bool {
	return hasLetter(remark) && strings.ToUpper(remark) == remark
}

func hasLetter(remark string) bool {
	matches, _ := regexp.MatchString("[a-zA-Z]", remark)
	return matches
}

func isForcefulQuestion(remark string) bool {
	return isQuestion(remark) && isShouting(remark)
}

func isAskingAQuestion(remark string) bool {
	return !isShouting(remark) && isQuestion(remark)
}

func isSilence(remark string) bool {
	return strings.TrimSpace(remark) == ""
}
