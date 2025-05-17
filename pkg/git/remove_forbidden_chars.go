package git

import "strings"

// List of characters that are not allowed in branch names.
// The list is not stricly follow the guide, just aim to remove problematic chars
//
// @see https://git-scm.com/docs/git-check-ref-format
var charsForbiddenInBranchName = []string{
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
	"'", "\"", "`",
	"«", "»",
	"!",
	"?",
	"#",
	"$",
	"*",
	"\\",
	"~",
	"^",
	"+",
	"=",
	"|",
}

func removeForbiddenChars(input string, connectorChar string) string {
	result := input
	// replace accent chars with special desired collections
	for _, r := range charsForbiddenInBranchName {
		result = strings.ReplaceAll(result, r, connectorChar)
	}
	return result
}
