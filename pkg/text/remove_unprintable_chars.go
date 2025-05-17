package text

import (
	"strings"
	"unicode"
)

// Remove inprintable chars
func RemoveUnprintableChars(input string) string {
	result := strings.Map(func(r rune) rune {
		if unicode.IsPrint(r) {
			return r
		}
		return -1
	}, input)
	return result
}
