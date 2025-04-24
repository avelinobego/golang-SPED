package utils

import "unicode"

func UnCapitalize(text string) string {
	result := make([]byte, 0)
	result = append(result, byte(unicode.ToLower(rune(text[0]))))
	result = append(result, text[1:]...)
	return string(result)

}
