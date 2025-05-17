package text

// Remove all consequent characters of provided char to one from given string
func RemoveConsequentChars(input string, char string) string {
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
