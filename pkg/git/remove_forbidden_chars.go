package git

import "strings"

// List of characters that are not allowed (or not preferred) in branch names.
// The list is not stricly following the guide, just aim to remove most
// problematic chars for general multi-language usage.
//
// @see https://git-scm.com/docs/git-check-ref-format
var charsForbiddenInBranchName = []string{
	"\n", "\r", "\t",
	"-", "_",
	":",
	";",
	",",
	"@",
	"(", ")",
	"{", "}",
	"[", "]",
	"<", ">",
	"'", "\"", "`", "’",
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
	"#",
	// custom
	"—", // em dash
	"…",
	// & is special char, see ReplaceCommonChars
	"•",
	"：",
}

var charsAllowedInBranchName = []string{
	".",
	"/",
}

func removeForbiddenChars(input string, connectorChar string) string {
	result := input
	// replace accent chars with special desired collections
	for _, r := range charsForbiddenInBranchName {
		result = strings.ReplaceAll(result, r, connectorChar)
	}
	return result
}

func trimAllowedChars(input string) string {
	result := input
	for _, v := range charsAllowedInBranchName {
		result = strings.Trim(result, v)
	}
	return result
}
