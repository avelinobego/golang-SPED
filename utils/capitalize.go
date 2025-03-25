package utils

import (
	"log"
	"unicode"
	"unicode/utf8"
)

func Capitalize(text string) string {

	r, size := utf8.DecodeRuneInString(text)
	if r == utf8.RuneError {
		log.Printf(`erro ao capitalizar: "%v"`, r)
		return text
	}
	return string(unicode.ToUpper(r)) + text[size:]

}
