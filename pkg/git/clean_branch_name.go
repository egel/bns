package git

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/egel/bns/pkg/text"
)

// Clean the given list of arguments/strings from unwanted or unprintable chars
func CleanBranchName(
	inputs []string,
	connectorChar string,
	keepOriginalCase bool,
	forceAscii bool,
) string {
	fmt.Println(inputs, connectorChar, keepOriginalCase, forceAscii)

	oneBuf := new(bytes.Buffer)
	for i, v := range inputs {
		tmp := v
		if !keepOriginalCase {
			tmp = strings.ToLower(tmp)
		}
		tmp = strings.Trim(tmp, " ")
		tmp = strings.ReplaceAll(tmp, " ", connectorChar)
		tmp = text.ReplaceCommonChars(tmp)
		tmp = text.RemoveUnprintableChars(tmp)
		tmp = removeForbiddenChars(tmp, connectorChar)

		if forceAscii {
			tmp = text.ForceToASCII(tmp)
		} else {
			tmp = text.CleanAccentChars(tmp)
		}
		tmp = trimAllowedChars(tmp)
		if i == 0 {
			oneBuf.WriteString(tmp)
		} else {
			oneBuf.WriteString(connectorChar + tmp)
		}
	}
	oneStr := oneBuf.String()
	oneStr = text.RemoveConsequentChars(oneStr, connectorChar)
	oneStr = strings.Trim(oneStr, connectorChar)
	return oneStr
}
