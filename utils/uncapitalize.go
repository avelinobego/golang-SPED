package utils

import (
	"log"
	"unicode"
	"unicode/utf8"
)

func UnCapitalize(text string) string {

	r, size := utf8.DecodeRuneInString(text)
	if r == utf8.RuneError {
		log.Printf(`erro ao decapitalizar: "%v"`, r)
		return text
	}
	return string(unicode.ToLower(r)) + text[size:]

}
