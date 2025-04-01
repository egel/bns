package git

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCleanBranchName(t *testing.T) {
	type testStrings struct {
		original []string
		expect   string
	}

	var tests = []testStrings{
		// German panagram
		{
			original: []string{"»Fix, Schwyz!«, quäkt Jürgen blöd vom Paß."},
			expect:   "fix-schwyz-quaekt-juergen-bloed-vom-pass",
		},

		// Polish panagram
		{
			original: []string{"Zażółć gęślą jaźń"},
			expect:   "zazolc-gesla-jazn",
		},

		// Funky panagram
		{
			original: []string{"Tĥïŝ ĩš â fůňķŷ Šťŕĭńġ"},
			expect:   "this-is-a-funky-string",
		},

		// single arguments
		{
			original: []string{" asdf  1    2 3   "},
			expect:   "asdf-1-2-3",
		},
		{
			original: []string{"cool & still"},
			expect:   "cool-and-still",
		},

		// multiple arguments
		{
			original: []string{"asdf", " 1", "2 ", " 3 "},
			expect:   "asdf-1-2-3",
		},
		{
			original: []string{"asdf", "  1", "2  ", "   3   "},
			expect:   "asdf-1-2-3",
		},
		{
			original: []string{"new member   ", "  1"},
			expect:   "new-member-1",
		},
	}

	for _, test := range tests {
		result := CleanBranchName(test.original)
		assert.Equal(t, test.expect, result)
	}
}
