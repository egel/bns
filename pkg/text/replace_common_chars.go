package text

import "fmt"

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
func ReplaceCommonChars(input string) string {
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
