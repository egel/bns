package text

// Remove all characters that are not in ASCII list (32-127)
func ForceToASCII(input string) string {
	result := make([]rune, 0, len(input))
	for _, r := range input {
		if r <= 127 {
			result = append(result, r)
		}
	}
	return string(result)
}
