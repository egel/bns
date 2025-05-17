package git

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// Default character use for connecting substrings
const DefaultConnectorChar = "-"

// List of characters to be removed
var charsToBeRemoved = []string{
	"\n", "\r", "\t",
	":",
	";",
	",",
	".",
	"@",
	"(", ")",
	"{", "}",
	"[", "]",
	"<", ">",
	"'", "\"",
	"«", "»",
	"!",
	"?",
}

// Clean the given list of arguments/strings from unwanted or unprintable chars
func CleanBranchName(inputs []string) string {
	oneBuf := new(bytes.Buffer)
	for i, j := range inputs {
		tmp := strings.ToLower(j)
		tmp = strings.Trim(tmp, " ") // each end
		tmp = strings.ReplaceAll(tmp, " ", DefaultConnectorChar)

		// replace accent chars with special desired collections
		for _, r := range charsToBeRemoved {
			tmp = strings.ReplaceAll(tmp, r, DefaultConnectorChar)
		}

		tmp = replaceKnownChars(tmp)
		tmp = cleanAccentChars(tmp)

		// remove inprintable chars (e.g. all unicode)
		tmp = strings.Map(func(r rune) rune {
			if unicode.IsPrint(r) {
				return r
			}
			return -1
		}, tmp)

		if i == 0 {
			oneBuf.WriteString(tmp)
		} else {
			oneBuf.WriteString(DefaultConnectorChar + tmp)
		}
	}

	oneStr := oneBuf.String()
	oneStr = removeCharsDuplicates(oneStr, DefaultConnectorChar)
	oneStr = strings.Trim(oneStr, DefaultConnectorChar)
	return oneStr
}

// Remove all consequent characters of provided char from given string
func removeCharsDuplicates(input string, char string) string {
	newStr := ""
	for _, c := range input {
		if string(c) == char && len(newStr) > 0 {
			lastChar := newStr[len(newStr)-1:]
			if lastChar != char {
				newStr += string(c)
			}
		} else {
			newStr += string(c)
		}
	}

	return newStr
}

// List of known characters to be swaped
var knownCharsMapping [][]string = [][]string{
	// Conventional
	{"&", fmt.Sprintf("%sand%s", DefaultConnectorChar, DefaultConnectorChar)},
	// German
	{"ä", "ae"},
	{"ü", "ue"},
	{"ö", "oe"},
	{"ß", "ss"},
	// Polish
	{"ł", "l"},
}

// Replace all chars from given test with ones from knownCharsMapping
func replaceKnownChars(input string) string {
	result := ""
	for _, c := range input {
		for i, k := range knownCharsMapping {
			if string(c) == k[0] {
				result += k[1]
				break
			}
			if i == len(knownCharsMapping)-1 {
				result += string(c)
			}
		}
	}
	return result
}

// Remove accent chars from given text
func cleanAccentChars(input string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	s, _, _ := transform.String(t, input)
	return s
}

// Remove all characters that are not in ASCII list
func forceASCII(input string) string {
	result := make([]rune, 0, len(input))
	for _, r := range input {
		if r <= 127 {
			result = append(result, r)
		}
	}
	return string(result)
}
