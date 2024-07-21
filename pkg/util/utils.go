package util

import (
	"regexp"
	"strings"
)

func StripTitle(title string, removeWhitespace bool, removePartsAfterColon bool) string {
	strippedTitle := title
	if removeWhitespace {
		strippedTitle = strings.Replace(title, " ", "", -1)
	}

	//Remove series name in parentheses, if it exists
	par1 := strings.Index(strippedTitle, "(")
	par2 := strings.Index(strippedTitle, ")")

	if par1 > -1 && par2 > -1 {
		strippedTitle = strippedTitle[0:par1] + strippedTitle[par2+1:]
		//Remove potential multiple whitespaces
		strippedTitle = strings.Join(strings.Fields(strippedTitle), " ")
	}

	if removePartsAfterColon && strings.Index(strippedTitle, ": ") > -1 {
		strippedTitle = strings.Split(strippedTitle, ": ")[0]
	}

	strippedTitle = regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(strippedTitle, "")
	strippedTitle = strings.ToLower(strippedTitle)
	return strings.TrimSpace(strippedTitle)
}
