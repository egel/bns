package display

import "fmt"

func ShowResults(original []string, updated string) {
	fmt.Printf("Original: %v\n\n", original)
	fmt.Printf("Updated: %s\n", updated)
}
