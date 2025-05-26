package git

import (
	"testing"

	"github.com/egel/bns/pkg/text"
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

		// Czech panagram
		{
			original: []string{"Příliš žluťoučký kůň úpěl ďábelské ódy"},
			expect:   "prilis-zlutoucky-kun-upel-dabelske-ody",
		},

		// French panagram
		{
			original: []string{"Dès Noël où un zéphyr haï me vêt de glaçons würmiens je dîne d’exquis rôtis de bœuf au kir à l’aÿ d’âge mûr & cætera"},
			expect:   "des-noel-ou-un-zephyr-hai-me-vet-de-glacons-wuermiens-je-dine-d-exquis-rotis-de-boeuf-au-kir-a-l-ay-d-age-mur-and-caetera",
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
		result := CleanBranchName(test.original, "-", false, false)
		assert.Equal(t, test.expect, result)
	}
}

func Test_CleanBranchName_WithCustomConnector(t *testing.T) {
	type testStrings struct {
		original []string
		expect   string
	}

	var tests = []testStrings{
		{
			original: []string{"Develop AI-powered bug reporter"},
			expect:   "develop_ai_powered_bug_reporter",
		},
	}

	for _, test := range tests {
		result := CleanBranchName(test.original, "_", false, false)
		assert.Equal(t, test.expect, result)
	}
}

func Test_CleanBranchName_WithKeepCase(t *testing.T) {
	type testStrings struct {
		original []string
		expect   string
	}

	var tests = []testStrings{
		{
			original: []string{"ISS-1234: Develop AI-powered bug reporter"},
			expect:   "ISS-1234-Develop-AI-powered-bug-reporter",
		},
	}

	for _, test := range tests {
		result := CleanBranchName(test.original, text.DefaultConnectorChar, true, false)
		assert.Equal(t, test.expect, result)
	}
}

func Test_CleanBranchName_WithForceAscii(t *testing.T) {
	type testStrings struct {
		original []string
		expect   string
	}

	var tests = []testStrings{
		// Chinese (Mandarian)
		{
			original: []string{"BUG：login功能缺陷"},
			expect:   "bug-login",
		},
		// Arabic
		{
			original: []string{"مشكلة: غيابية في fonctionnalité d'inscription"},
			expect:   "fonctionnalit-d-inscription",
		},
	}

	for _, test := range tests {
		result := CleanBranchName(test.original, text.DefaultConnectorChar, false, true)
		assert.Equal(t, test.expect, result)
	}
}
