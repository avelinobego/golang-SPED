package utils

import "strings"

func CleanBetween(initial string, start string, end string) string {
	result := strings.Builder{}
	include := true
	for _, v := range initial {
		if string(v) == start {
			include = false
		}
		if include {
			result.WriteRune(v)
		}
		if string(v) == end {
			include = true
		}
	}

	return result.String()
}
