// Package acronym provides a function to create acronyms (or initialisms) from
// strings.
package acronym

import (
	"regexp"
	"strings"
)

const delimiterPattern = `[\s\-]`

// Abbreviate converts a string to its acronym representation.
// For example, Abbreviate("Man in the middle") == "MITM"
func Abbreviate(s string) string {
	var result string
	re := regexp.MustCompile(delimiterPattern)
	for _, word := range re.Split(s, -1) {
		result = result + strings.ToUpper(string(word[0]))
	}
	return result
}
