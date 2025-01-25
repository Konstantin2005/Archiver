package lib

import (
	"strings"
	"unicode"
)

func Encode(str string) string {

	return ""
}

// prepareText prepare text to be fit for encode :
// changes upper case бла бла бла : 1 + lower case letter
// i.g: My name is Ted -> !my name is !Ted
func prepareText(str string) string {

	var duf strings.Builder

	for _, ch := range str {
		if unicode.IsUpper(ch) {
			duf.WriteRune('!')
			duf.WriteRune(unicode.ToLower(ch))
		} else {
			duf.WriteRune(ch)
		}
	}
	return duf.String()
}
