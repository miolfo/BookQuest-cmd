package util

import (
	"regexp"
	"strings"
)

func StripTitle(title string, removeWhitespace bool) string {
	strippedTitle := title
	if removeWhitespace {
		strippedTitle = strings.Replace(title, " ", "", -1)
	}

	//Remove series name in parentheses, if it exists
	par1 := strings.Index(title, "(")
	par2 := strings.Index(title, ")")

	if par1 > 0 && par2 > 0 {
		strippedTitle = strippedTitle[0 : par1-1]
	}

	strippedTitle = regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(strippedTitle, "")
	strippedTitle = strings.ToLower(strippedTitle)
	return strippedTitle
}
